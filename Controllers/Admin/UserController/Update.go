package UserController

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
	"github.com/vahidlotfi71/ticket/Validations"
)

func Update(c *fiber.Ctx) error {
	// دریافت ID از پارامتر URL
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"message": "شناسه کاربر الزامی است"})
	}

	// دریافت داده‌های جدید
	var user Models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ورودی نامعتبر"})
	}

	// اجرای Validation
	v := Validations.UserValidation(user)
	if !v.IsValid() {
		return c.Status(422).JSON(fiber.Map{"errors": v.Errors()})
	}

	// تبدیل ID از string به uint
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "شناسه کاربر نامعتبر است"})
	}
	user.ID = uint(uid)

	// بروزرسانی مستقیم در دیتابیس
	if err := Config.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "کاربر با موفقیت بروزرسانی شد",
		"data":    user,
	})
}
