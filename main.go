package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

	var jsonResponse string
	if valid {
		fmt.Println("Received Input from Database:")
		fmt.Print("\n")
		fmt.Println(users)
		fmt.Print("\n")
		fmt.Println("---------------------------------------------")
		fmt.Print("\n")

		postURL := "https://6ir887qv2c.execute-api.us-east-2.amazonaws.com/test/userdataproxy"
		postBody, _ := json.Marshal(users)
		resp, err := http.Post(postURL, "application/json", bytes.NewBuffer(postBody))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		jsonResponse = string(body)

		fmt.Println("Http POST response: ")
		fmt.Print("\n")
		fmt.Println(jsonResponse)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		var output string
		output += "Received Input from Database: \n"
		for _, user := range users {
			output += fmt.Sprintf("%+v", user) + "\n"
		}

		output += "\n" + "Http POST response: \n" + jsonResponse
		return c.SendString(output)
	})

	app.Listen(":3000")
}
