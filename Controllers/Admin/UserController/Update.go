package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
)

func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var user Models.User

	// پیدا کردن کاربر
	if err := Config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// گرفتن داده‌های جدید
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// ذخیره تغییرات
	if err := Config.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error updating user",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    user,
	})
}
