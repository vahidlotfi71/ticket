package UserController

import (
	"vahid/Config"
	"vahid/Models"

	"github.com/gofiber/fiber/v2"
)

func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var user Models.Users

	if err := Config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": user,
	})
}
