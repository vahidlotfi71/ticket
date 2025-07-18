package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
)

func Store(c *fiber.Ctx) error {
	user := new(Models.User)

	// گرفتن داده‌های ورودی
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// ذخیره در دیتابیس
	if err := Config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error saving user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user,
	})
}
