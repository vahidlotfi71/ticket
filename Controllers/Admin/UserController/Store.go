package UserController

import (
	"vahid/Config"
	"vahid/Models"

	"github.com/gofiber/fiber/v2"
)

func Store(c *fiber.Ctx) error {
	var user Models.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input.",
		})
	}

	if err := Config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully.",
	})

}
