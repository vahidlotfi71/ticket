package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
)

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var user Models.User

	// پیدا کردن کاربر
	if err := Config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// حذف کاربر
	if err := Config.DB.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error deleting user",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
