package main

import (
	eh_system "eh_system/lib"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

const IMAGE_DATA_PATH string="C:\\Users\\ktkm\\Desktop\\h\\cg"
const THUMBNAIL_DATA_PATH string="C:\\Users\\ktkm\\Desktop\\eh-system3\\thumbnaildatas\\thumbnaildata"

func main() {
    // --- variables ---
    var HERE string
    HERE,_=os.Executable()


    // --- server setup ---
	var app *fiber.App = fiber.New(fiber.Config {
        CaseSensitive: true,
        EnablePrintRoutes: true,
    })



    // ---- apis ----
    // get a single album for viewing mode
    app.Post("/get-album",func (ctx *fiber.Ctx) error {
        var targetPath string=string(ctx.Body())

        fmt.Println("getting album:",targetPath)

        var result []string=eh_system.GetAllImagesFlat(IMAGE_DATA_PATH,targetPath,true)

        fmt.Printf("sending %d images",len(result))

        return ctx.JSON(eh_system.AlbumResponse {
            Urls:result,
            Mode:eh_system.LOCAL_MODE,
        })
    })

    // get a list of album infos under a target path
    app.Post("/get-album-info",func(ctx *fiber.Ctx) error {
        var targetPath string=string(ctx.Body())

        fmt.Println("getting album info:",targetPath)

        var result []eh_system.AlbumInfo=eh_system.GetAlbumInfos(IMAGE_DATA_PATH,targetPath)

        fmt.Printf("sending %d albums\n",len(result))

        return ctx.JSON(result)
    })



    // --- statics ---
    app.Static("/build",filepath.Join(HERE,"eh-system-web/build"))

    app.Static("/viewer/*",filepath.Join(HERE,"eh-system-web/web/pages/ehviewer"))

    app.Static("/albums/*",filepath.Join(HERE,"eh-system-web/web/pages/albumexplore"))

    app.Static("/assets/fonts",filepath.Join(HERE,"eh-system-web/web/assets/fonts"))

    app.Static("/assets/imgs",filepath.Join(HERE,"eh-system-web/web/assets/imgs"))

    app.Static("/imagedata",IMAGE_DATA_PATH,fiber.Static{
        Browse:true,
    })

    app.Static("/thumbnaildata",THUMBNAIL_DATA_PATH)


    // --- serve ---
    app.Listen("localhost:4200")
}