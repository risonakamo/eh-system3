package main

import (
	eh_system "eh_system/lib"

	"github.com/davecgh/go-spew/spew"
	// "github.com/gofiber/fiber/v2"
)

func main() {
    // fmt.Println("bruh")
    // var app *fiber.App=fiber.New()

    // app.Get("/",func (c *fiber.Ctx) error{
    // 	return c.SendString("asdasd")
    // })

    // app.Listen(":4200")

    var result=eh_system.GetAllImages(
        "C:\\Users\\ktkm\\Desktop\\h\\cg",
        "nekonote",
    )

    // fmt.Println(result)
    spew.Dump(result)
}