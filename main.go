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

	//Opens database connection and returns *sql.DB for future database functions
	db := database.ConnectDB()
	defer db.Close()

	//The unique user_id's used to query the database
	ids := []int{1, 2, 3}

	//Function returns array of User structs after querying database based on the unique user_id's
	users, err := database.QueryDBforUsers(db, ids)
	if err != nil {
		panic(err)
	}

	//Validation to check if any fields are empty in retrieved user data
	valid, err := data.ValidateUsers(users)
	if err != nil {
		panic(err)
	}

	var jsonResponse string
	if valid {
		//Printing retrieved user data to STDOUT
		fmt.Println("Received Input from Database:")
		fmt.Print("\n")
		fmt.Println(users)
		fmt.Print("\n")
		fmt.Println("---------------------------------------------")
		fmt.Print("\n")

		//Forming request to API Gateway Endpoint with retrieved user data in request body
		postURL := "https://6ir887qv2c.execute-api.us-east-2.amazonaws.com/test/userdataproxy"
		postBody, _ := json.Marshal(users)
		resp, err := http.Post(postURL, "application/json", bytes.NewBuffer(postBody))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		//Decode response from API Gateway and print to STDOUT
		body, err := io.ReadAll(resp.Body)
		jsonResponse = string(body)

		fmt.Println("Http POST response: ")
		fmt.Print("\n")
		fmt.Println(jsonResponse)
	}

	//Was checking out Go Fiber, so the request artefacts can be viewed at localhost:3000 too
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
