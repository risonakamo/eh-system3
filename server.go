package main

import (
	"fmt"
	"io/fs"

	"os"
	"path/filepath"

	"facette.io/natsort"
	"github.com/davecgh/go-spew/spew"
	// "github.com/gofiber/fiber/v2"
)

/* a top level album item. an album can recursively contain other albums, or be a leaf album, only
   containing images. both are represented by this struct */
type AlbumInfo struct {
	title string
	// total number of items in the album (recursive)
	items int
	// items at the top level of the album
	immediateItems int

	// path to random img in the album
	img string
	// last modified date of the album
	date string
	// if this is a leaf album or not (only contains images and no other folders)
	album bool
}

func main() {
	// fmt.Println("bruh")
	// var app *fiber.App=fiber.New()

	// app.Get("/",func (c *fiber.Ctx) error{
	// 	return c.SendString("asdasd")
	// })

	// app.Listen(":4200")

	var result=getAllImages(
		"C:\\Users\\ktkm\\Desktop\\h\\cg",
		"nekonote",
	)

	// fmt.Println(result)
	spew.Dump(result)
}

func getAlbumInfo(imageDataPath string,targetPath string) {
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
func getAllImages(imageDataPath string,targetPath string) [][]string {
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

		var result [][]string=getAllImages(imageDataPath,dir)

		collectedItems=append(collectedItems,result...)
	}

	// adding own dir's items to subdir items (if any)
	if len(imagesInCurrentDir)>0 {
		natsort.Sort(imagesInCurrentDir)
		collectedItems=append(collectedItems,imagesInCurrentDir)
	}

	return collectedItems
}