package main

import (
	"github.com/davecgh/go-spew/spew"
	// "github.com/gofiber/fiber/v2"

	"eh_system/album_info"
)

func main() {
	// fmt.Println("bruh")
	// var app *fiber.App=fiber.New()

	// app.Get("/",func (c *fiber.Ctx) error{
	// 	return c.SendString("asdasd")
	// })

	// app.Listen(":4200")

	var result=album_info.GetAllImages(
		"C:\\Users\\ktkm\\Desktop\\h\\cg",
		"nekonote",
	)

	// fmt.Println(result)
	spew.Dump(result)
}