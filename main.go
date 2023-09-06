package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanProd/FrameplayTakehome/database"
	"github.com/ryanProd/FrameplayTakehome/jsonUtil"
)

func main() {
	app := fiber.New()

	db := database.ConnectDB()
	defer db.Close()

	ids := []int{1, 2, 3}

	users, err := database.QueryDBforUsers(db, ids)
	if err != nil {
		panic(err)
	}

	for _, val := range users {
		fmt.Println(val.Username)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(jsonUtil.UploadJson("Frameplay"))
	})

	app.Listen(":3000")
}
