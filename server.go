package main

import (
	"fmt"
	"io/fs"

	"os"
	"path/filepath"
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

	getAlbumInfo(
		"C:\\Users\\ktkm\\Desktop\\h\\cg",
		"nekonote",
	)
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