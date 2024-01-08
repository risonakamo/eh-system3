package eh_system

import (
	"io/fs"
	"os"
	"path/filepath"
)

// return paths of images under target image dir that do not have thumbnails
// in the corresponding thumbnail dir
func FindMissingImagesWithoutThumbnails(
    imageDir string,
    thumbnailDir string,
) []string {
    var res []string

    filepath.Walk(imageDir,func (path string,info fs.FileInfo,err error) error {
        // if found a thumbnailable file while scanning the target folder
        if info.IsDir() || !thumbnailableFile(path) {
            return nil
        }

        // try to resolve its thumbnail path
        var thumbnailPath string=findThumbnailOfImage(
            imageDir,
            thumbnailDir,
            path,
        )

        // check if that thumbnail exists
        if !isFile(thumbnailPath) {
            res=append(res,thumbnailPath)
        }

        return nil
    })

    return res
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