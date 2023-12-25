// functions for getting things related to albums
package eh_system

import (
	"fmt"
	"io/fs"

	"os"
	"path/filepath"

	"facette.io/natsort"
)

func GetAlbumInfo(imageDataPath string,targetPath string) {
    var fullTargetPath string=filepath.Join(imageDataPath,targetPath)

    var files []fs.DirEntry
    var err error
    files,err=os.ReadDir(fullTargetPath)

    if err!=nil {
        panic("failed to readdir")
    }

    fmt.Println(files)
}

/* get ALL images under target album, recursively.
   paths are relative to the image data path (will not include the full file path).
   return images are grouped by their albums, which have an ordering. */
func GetAllImages(imageDataPath string,targetPath string) [][]string {
    var fullTargetPath string=filepath.Join(imageDataPath,targetPath)

    var direntrys []fs.DirEntry
    var err error
    direntrys,err=os.ReadDir(fullTargetPath)

    if err!=nil {
        panic("read dir failed")
    }

    var imagesInCurrentDir []string
    var dirsInCurrentDir []string

    // sort all items in current dir into image or dir
    for i := range direntrys {
        var direntry fs.DirEntry=direntrys[i]

        var itemPath string=filepath.ToSlash(filepath.Join(targetPath,direntry.Name()))

        if direntry.IsDir() {
            dirsInCurrentDir=append(
                dirsInCurrentDir,
                itemPath,
            )
        } else {
            imagesInCurrentDir=append(imagesInCurrentDir,itemPath)
        }
    }

    var collectedItems [][]string

    // recursively collecting items of sub dirs
    for i := range dirsInCurrentDir {
        var dir string=dirsInCurrentDir[i]

        var result [][]string=GetAllImages(imageDataPath,dir)

        collectedItems=append(collectedItems,result...)
    }

    // adding own dir's items to subdir items (if any)
    if len(imagesInCurrentDir)>0 {
        natsort.Sort(imagesInCurrentDir)
        collectedItems=append(collectedItems,imagesInCurrentDir)
    }

    return collectedItems
}