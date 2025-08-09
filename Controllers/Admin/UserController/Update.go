package UserController

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
	rules "github.com/vahidlotfi71/ticket/Rules"
)

func Update(c *fiber.Ctx) error {
	/* Taking the id URL paramter to be able to fetch the user */
	id := c.Params("id")
	var user Models.User
	Config.DB.First(&user, id)
	/* Check if user exists in the database */
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Check if the request body is a valid json object
	if !json.Valid(c.Body()) {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid json encoding",
		})
	}

	var newUser Models.User
	// Then trying to unmarshal/parse the json object
	if err := json.Unmarshal(c.Body(), &newUser); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "could not parse the json",
		})
	}

	validator := rules.NewValidator()
	if newUser.Name != "" {
		validator.SetField("name", newUser.Name).
			MinLength(3, "Name must be at least 3 characters").
			MaxLength(255, "Name must not exceed 255 characters")
	}

	if newUser.Email != "" {
		validator.SetField("email", newUser.Email).
			IsEmail("The email format is invalid").
			MaxLength(255, "Email must not exceed 255 characters")
	}

	if newUser.Password != "" {
		validator.SetField("password", newUser.Password).
			MinLength(3, "Password must be at least 6 characters long").
			MaxLength(255, "Password must not exceed 255 characters")
	}

	if newUser.Phone != "" {
		validator.SetField("phone", newUser.Phone).
			ExactLength(11, "Phone number must contain exactly 11 digits").
			StartsWith("09", "The phone number must start with 09")
	}

	if err := validator.Validate(c); err != nil {
		return err
	}

	if newUser.Name != "" {
		user.Name = newUser.Name
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Password != "" {
		user.Password = newUser.Password
	}

	// There is no need to specify the table, because we are
	// using a model instance to create the record, so the gorm
	// will figure it out by itself to handle it in the relative
	// table in database
	if err := Config.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "user updated successfully",
	})
	/* if we want to handle it manually, we can use .Exec() method */
}
