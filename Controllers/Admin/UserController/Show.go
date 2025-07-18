package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
)

func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var user Models.User

	// پیدا کردن کاربر
	if err := Config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": user,
	})
}
