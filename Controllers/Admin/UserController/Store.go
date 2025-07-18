package UserController

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
)

func Store(c *fiber.Ctx) error {
	// Check if the request body is a valid json object
	if !json.Valid(c.Body()) {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid json encoding",
		})
	}
	// Then trying to unmarshal/parse the json object
	var data Models.User
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "could not parse the json",
		})
	}
	// There is no need to specify the table, because we are
	// using a model instance to create the record, so the gorm
	// will figure it out by itself to handle it in the relative
	// table in database
	if err := Config.DB.Create(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "user created successfully",
	})
	/* if we want to handle it manually, we can use .Exec() method */
}
