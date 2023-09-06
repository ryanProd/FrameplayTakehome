package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ryanProd/FrameplayTakehome/jsonUtil"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(jsonUtil.UploadJson("Frameplay"))
	})

	app.Listen(":3000")
}
