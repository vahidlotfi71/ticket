package UserController

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
	rules "github.com/vahidlotfi71/ticket/Rules"
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

	validator := rules.NewValidator()
	// اعتبار سنجی فیلد ها
	validator.SetField("name", data.Name).Required("The name is required").
		MinLength(3, "Name must be at least 3 characters").
		MaxLength(255, "Name must not exceed 255 characters")

	validator.SetField("email", data.Email).Required("The email is required").
		IsEmail("The email format is invalid").
		MaxLength(255, "Email must not exceed 255 characters")

	validator.SetField("password", data.Password).
		Required("The password is required").
		MinLength(6, "Password must be at least 6 characters long").
		MaxLength(255, "Password must not exceed 255 characters")

	validator.SetField("phone", data.Phone).
		Required("The phone number is required").
		ExactLength(11, "Phone number must contain exactly 11 digits").
		StartsWith("09", "The phone number must start with 09")

	if err := validator.Validate(c); err != nil {
		return err
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
