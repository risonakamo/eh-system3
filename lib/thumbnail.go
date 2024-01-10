package eh_system

import (
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

// return paths of images under target image dir that do not have thumbnails
// in the corresponding thumbnail dir
func FindMissingImagesWithoutThumbnails(
    imageDir string,
    thumbnailDir string,
    workers int,
) []string {
    var res []string
    var imagePaths []string

    filepath.WalkDir(imageDir,func(path string,info fs.DirEntry,err error) error {
        if info.IsDir() || !thumbnailableFile(path) {
            return nil
        }

        imagePaths=append(imagePaths,path)

        return nil
    })

    var wg sync.WaitGroup
    var imagePathsChannel chan string=make(chan string)
    var missingThumbnailsChannel chan string=make(chan string)

    // spawn workers
    for i:=0;i<workers;i++ {
        wg.Add(1)
        go thumbnailCheckWorker(
            imagePathsChannel,
            missingThumbnailsChannel,
            &wg,

            imageDir,
            thumbnailDir,
        )
    }

    // submit jobs
    for i:=range imagePaths {
        imagePathsChannel<-imagePaths[i]
    }


    wg.Wait()

    return res
}

// reads img files from a list. converts them into thumbnail paths, and checks if they exist.
// if it doesn't, dumps the result into nonExistingThumbnails channel
func thumbnailCheckWorker(
    imagesChannel <-chan string,
    missingThumbnailsChannel chan<- string,
    wg *sync.WaitGroup,

    imageDir string,
    thumbnailDir string,
) {
    for {
        var imagePath string
        var ok bool
        imagePath,ok=<-imagesChannel

        if !ok {
            wg.Done()
            return
        }

        // try to resolve its thumbnail path
        var thumbnailPath string=findThumbnailOfImage(
            imageDir,
            thumbnailDir,
            imagePath,
        )

        // check if that thumbnail exists
        if !isFile(thumbnailPath) {
            missingThumbnailsChannel<-thumbnailPath
        }
    }
}

// return if the target file is thumbnailable
func thumbnailableFile(path string) bool {
    switch filepath.Ext(path) {
    case ".png",".jpg",".gif",".mp4":
        return true
    default:
        return false
    }
}

// checks if target is a file and it exists
func isFile(path string) bool {
    var stat fs.FileInfo
    var err error
    stat,err=os.Stat(path)

    if err!=nil {
        return false
    }

    return !stat.IsDir()
}