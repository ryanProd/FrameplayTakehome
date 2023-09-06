package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/ryanProd/FrameplayTakehome/config"
	"github.com/ryanProd/FrameplayTakehome/jsonUtil"
)

func main() {
	app := fiber.New()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"), config.Config("DB_HOST"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var version string
	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}

	fmt.Printf("version=%s\n", version)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(jsonUtil.UploadJson("Frameplay"))
	})

	app.Listen(":3000")
}
