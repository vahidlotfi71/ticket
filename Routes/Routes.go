package routes

import (
	"vahid/Controllers/Admin/UserController"

	"github.com/gofiber/fiber/v2"
)

var defaultController = func(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "hello from default controller",
	})
}

func Routes(app *fiber.App) {
	adminGroup := app.Group("/admin")

	adminGroup.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"mesage": "admin index route",
		})
	})

	userGroup := adminGroup.Group("/user")

	userGroup.Get("/show/:id", UserController.Show)
	userGroup.Post("/store", UserController.Store)

	/* Static file rendering */
	app.Static("/", "public")

	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Route Not found",
		})
	})
}
