// functions for getting things related to albums
package eh_system

import (
	"io/fs"

	"os"
	"path/filepath"

	"facette.io/natsort"
)

/** get album info for every item in a target path */
func GetAlbumInfos(imageDataPath string,targetPath string) []AlbumInfo {
    var fullTargetPath string=filepath.Join(imageDataPath,targetPath)

    var items []fs.DirEntry
    items,_=os.ReadDir(fullTargetPath)

    var albums []AlbumInfo

    for i := range items {
        var item *fs.DirEntry=&items[i]

        if !(*item).IsDir() {
            continue
        }

        var subdirPath string=filepath.Join(targetPath,(*item).Name())

        albums=append(albums,getAlbumInfo(imageDataPath,subdirPath))
    }

    return albums
}

/** get album info of single target path */
func getAlbumInfo(imageDataPath string,targetPath string) AlbumInfo {
    var fullTargetPath string=filepath.Join(imageDataPath,targetPath)

    var allItems []string=GetAllImagesFlat(imageDataPath,targetPath,false)

    var immediateItems []fs.DirEntry
    immediateItems,_=os.ReadDir(fullTargetPath)

    // todo: maybe get the highest mod time of the immediate items instead of
    // the target path itself?
    var targetInfo os.FileInfo
    targetInfo,_=os.Stat(fullTargetPath)

    return AlbumInfo {
        Title:filepath.Base(targetPath),
        Items:len(allItems),
        ImmediateItems:len(immediateItems),

        Img:randFromArray[string](allItems),
        Date:targetInfo.ModTime().String(),

        // it is an image album if all immediate items are non-directories
        Album:isAllFiles(immediateItems),
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

        collectedItems=append(collectedItems,(result)...)
    }

    // adding own dir's items to subdir items (if any)
    if len(imagesInCurrentDir)>0 {
        natsort.Sort(imagesInCurrentDir)
        collectedItems=append(collectedItems,imagesInCurrentDir)
    }

    if shuffle {
        shuffleArray[[]string](collectedItems)
    }

    return collectedItems
}

/* same as get all images, but the result is flattened */
func GetAllImagesFlat(imageDataPath string,targetPath string,shuffle bool) []string {
    var res [][]string=GetAllImages(imageDataPath,targetPath,shuffle)

    var output []string

    // flatten
    for i := range res {
        output=append(output,res[i]...)
    }

    return output
}