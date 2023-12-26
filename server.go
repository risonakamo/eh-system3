package main

import (
	eh_system "eh_system/lib"
	"fmt"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

const IMAGE_DATA_PATH string="C:\\Users\\ktkm\\Desktop\\h\\cg"

func main() {
	var app *fiber.App = fiber.New(fiber.Config {
        CaseSensitive: true,
        EnablePrintRoutes: true,
    })

    // ---- apis ----
    app.Post("/get-album",func (ctx *fiber.Ctx) error {
        var targetPath string=string(ctx.Body())

        fmt.Println("getting album:",targetPath)

        var result []string=eh_system.GetAllImagesFlat(IMAGE_DATA_PATH,targetPath,true)

        return ctx.JSON(eh_system.AlbumResponse {
            Urls:result,
            Mode:eh_system.LOCAL_MODE,
        })
    })

    // app.Use(logger.New())
    app.Listen(":4200")
}