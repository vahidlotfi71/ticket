package Middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
	"github.com/vahidlotfi71/ticket/Utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tokenSplit := strings.Split(token, " ")
	if len(tokenSplit) != 2 || strings.ToLower(tokenSplit[0]) != "bearer" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token type, expected bearer",
		})
	}

	user_id, name, phone, email, err := Utils.VerifyToken(tokenSplit[1])
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	tx := Config.DB
	var user Models.User

	if err := tx.Where("id = ? AND name = ? AND phone = ? AND email = ?",
		user_id, name, phone, email).
		First(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error retrieving user information from database",
		})
	}

	if user.ID == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("user", user)
	return c.Next()
}
