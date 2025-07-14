package Routes

import "github.com/gofiber/fiber/v3"

func Routes(app *fiber.App) {
	adminGroup := app.Group("admin")
	adminGroup.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "admin index route",
		})
	})
}
