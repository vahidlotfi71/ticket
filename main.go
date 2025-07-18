package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Commands"
	"github.com/vahidlotfi71/ticket/Config"
	routes "github.com/vahidlotfi71/ticket/Routes"
)

func main() {

	if err := Config.GetEnv(); err != nil {
		fmt.Printf("Error while loading the .env file :\n")
		fmt.Printf("\t%s\n", err.Error())
		os.Exit(2)
	} else {
		fmt.Printf("Loaded .env file successfully ...\n")
	}

	if err := Config.Connect(); err != nil {
		fmt.Printf("Error while connecting to the database :\n")
		fmt.Printf("\t%s\n", err.Error())
		os.Exit(2)
	} else {
		fmt.Printf("Connected to the database successfully ...\n")
	}

	Commands.AutoMigrate()

	app := *fiber.New()

	app.Get("/Hello", func(c *fiber.Ctx) error {
		return c.SendString("hello world ")
	})

	routes.Routes(&app)

	port := "8000"
	fmt.Printf("Starting the server on port %s\n", port)
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		fmt.Printf("Error while starting the server, see why :\n")
		fmt.Printf("\t%s\n", err.Error())
		os.Exit(2)
	}
	app.Listen(":8000")

}
