package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
	"github.com/vahidlotfi71/ticket/Validations"
)

func Store(c *fiber.Ctx) error {
	var user Models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input.",
		})
	}

	// استفاده از Chain Method Validation
	v := Validations.UserValidation(user)
	if !v.IsValid() {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": v.Errors(),
		})
	}

	// ذخیره در دیتابیس
	if err := Config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully.",
	})
}
