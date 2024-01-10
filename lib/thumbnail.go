package eh_system

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

// detect all missing thumbnails, and generate them all
func GenerateMissingThumbnails(
    imageDir string,
    thumbnailDir string,
    size int,

    workers int,
    suppressFfmpegError bool,
) {
    var missing []MissingThumbnail=FindMissingImagesWithoutThumbnails(
        imageDir,
        thumbnailDir,
        workers,
    )

    fmt.Printf("missing %v thumbnails\n",len(missing))

    var inputs []string
    var outputs []string

    for i := range missing {
        inputs=append(inputs,missing[i].srcItem)
        outputs=append(outputs,missing[i].neededThumbnail)
    }

    fmt.Println("generating...")

    GenThumbnails(
        inputs,
        outputs,
        size,
        workers,
        suppressFfmpegError,
    )
}

// return paths of images under target image dir that do not have thumbnails
// in the corresponding thumbnail dir
func FindMissingImagesWithoutThumbnails(
    imageDir string,
    thumbnailDir string,
    workers int,
) []MissingThumbnail {
    var res []MissingThumbnail
    var imagePaths []string

    // collect all valid image paths
    filepath.WalkDir(imageDir,func(path string,info fs.DirEntry,err error) error {
        if info.IsDir() || !thumbnailableFile(path) {
            return nil
        }

        imagePaths=append(imagePaths,path)

        return nil
    })

    var wg sync.WaitGroup
    var collectorWg sync.WaitGroup
    var imagePathsChannel chan string=make(chan string)
    var missingThumbnailsChannel chan MissingThumbnail=make(chan MissingThumbnail)

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

    // spawn worker to read from result channel
    collectorWg.Add(1)
    go func() {
        for {
            var missingThumbnail MissingThumbnail
            var ok bool
            missingThumbnail,ok=<-missingThumbnailsChannel

            if !ok {
                collectorWg.Done()
                return
            }

            res=append(res,missingThumbnail)
        }
    }()

    // submit jobs
    for i:=range imagePaths {
        imagePathsChannel<-imagePaths[i]
    }

    // all image paths sent, done
    close(imagePathsChannel)

    // wait for workers to complete last jobs
    wg.Wait()

    // once all workers done, close the result channel,
    // wait for collector channel to finish
    close(missingThumbnailsChannel)
    collectorWg.Wait()

    return res
}

// reads img files from a list. converts them into thumbnail paths, and checks if they exist.
// if it doesn't, dumps the result into nonExistingThumbnails channel
func thumbnailCheckWorker(
    imagesChannel <-chan string,
    missingThumbnailsChannel chan<- MissingThumbnail,
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
            missingThumbnailsChannel<-MissingThumbnail{
                srcItem:imagePath,
                neededThumbnail:thumbnailPath,
            }
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