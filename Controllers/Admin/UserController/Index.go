package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
)

func Index(c *fiber.Ctx) error {
	var users []Models.User
	tx := Config.DB.Model(&Models.User{})
	if err := tx.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": users,
	})

}
