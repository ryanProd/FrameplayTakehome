package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanProd/FrameplayTakehome/data"
	"github.com/ryanProd/FrameplayTakehome/database"
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

	valid, err := data.ValidateUsers(users)
	if err != nil {
		panic(err)
	}

	if valid {
		fmt.Print(users)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		var output string
		for _, user := range users {
			output += fmt.Sprintf("%+v", user) + "\n"
		}
		return c.SendString(output)
	})

	app.Listen(":3000")
}
