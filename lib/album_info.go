// functions for getting things related to albums
package eh_system

import (
	"io/fs"
	"math/rand"

	"os"
	"path/filepath"

	"facette.io/natsort"
)

// func GetAlbumInfo(imageDataPath string,targetPath string) {
//     var fullTargetPath string=filepath.Join(imageDataPath,targetPath)

//     var files []fs.DirEntry
//     var err error
//     files,err=os.ReadDir(fullTargetPath)

//     if err!=nil {
//         panic("failed to readdir")
//     }

//     fmt.Println(files)
// }

func GetAlbumInfo(imageDataPath string,targetPath string) AlbumInfo {
    var allItems []string=getAllImagesFlat(imageDataPath,targetPath,false)

    return AlbumInfo {
        title:filepath.Base(targetPath),
        items:len(allItems),
        immediateItems:0,

        img:"",
        date:"",
        album:false,
    }
}

/* get ALL images under target album, recursively.
   paths are relative to the image data path (will not include the full file path).
   return images are grouped by their albums, which have an ordering. */
func GetAllImages(imageDataPath string,targetPath string,shuffle bool) [][]string {
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

        var result [][]string=GetAllImages(imageDataPath,dir,shuffle)

        collectedItems=append(collectedItems,result...)
    }

    // adding own dir's items to subdir items (if any)
    if len(imagesInCurrentDir)>0 {
        natsort.Sort(imagesInCurrentDir)
        collectedItems=append(collectedItems,imagesInCurrentDir)
    }

    if shuffle {
        shuffleArray[[]string](&collectedItems)
    }

    return collectedItems
}

/* same as get all images, but the result is flattened */
func getAllImagesFlat(imageDataPath string,targetPath string,shuffle bool) []string {
    var res [][]string=GetAllImages(imageDataPath,targetPath,shuffle)

    var output []string

    // flatten
    for i := range res {
        output=append(output,res[i]...)
    }

    return output
}

/** shuffle an array */
func shuffleArray[T any](array *[]T) {
    rand.Shuffle(len(*array),func (i int,j int) {
        (*array)[i],(*array)[j]=(*array)[j],(*array)[i]
    })
}