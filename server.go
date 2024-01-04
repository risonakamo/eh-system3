package main

import (
	eh_system "eh_system/lib"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

// config path relative to exe
const CONFIG_PATH string="../config/config1.yml"
// const CONFIG_PATH string="../config/config2.yml"

func main() {
    // --- variables ---
    var HERE string
    HERE,_=os.Executable()
    HERE=filepath.Dir(HERE)

    // --- config get ---
    var config eh_system.EhSystemConfig=eh_system.LoadConfig(
        filepath.Join(HERE,CONFIG_PATH),
    )


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

        var result []string=eh_system.GetAllImagesFlat(config.ImageDir,targetPath,true)
        result=eh_system.FixImageUrls(result)

        fmt.Printf("-> sending %d images\n",len(result))

        return ctx.JSON(eh_system.AlbumResponse {
            Urls:result,
            Mode:eh_system.LOCAL_MODE,
        })
    })

    // get a list of album infos under a target path
    app.Post("/get-album-info",func(ctx *fiber.Ctx) error {
        var targetPath string=string(ctx.Body())

        var displayTargetPath string=targetPath

        if len(targetPath)==0 {
            displayTargetPath="/"
        }

        fmt.Println("getting album info:",displayTargetPath)

        var result []eh_system.AlbumInfo=eh_system.GetAlbumInfos(config.ImageDir,targetPath)
        result=eh_system.FixAlbumInfoImageUrls(result)

        fmt.Printf("-> sending %d albums\n",len(result))

        return ctx.JSON(result)
    })

    app.Use(redirect.New(redirect.Config{
        Rules:map[string]string {
            "/":"/albums",
        },
    }))



    // --- statics ---
    app.Static("/build",filepath.Join(HERE,"../eh-system-web/build"))

    app.Static("/viewer/*",filepath.Join(HERE,"../eh-system-web/web/pages/ehviewer"))

    app.Static("/albums/*",filepath.Join(HERE,"../eh-system-web/web/pages/albumexplore"))

    app.Static("/assets/fonts",filepath.Join(HERE,"../eh-system-web/web/assets/fonts"))

    app.Static("/assets/imgs",filepath.Join(HERE,"../eh-system-web/web/assets/imgs"))

    app.Static("/imagedata",config.ImageDir,fiber.Static{
        Browse:true,
    })

    app.Static("/thumbnaildata",config.ThumbnailDir)


    // --- serve ---
    app.Listen(":80")
}