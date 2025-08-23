package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Models"
	"github.com/vahidlotfi71/ticket/Validations"
)

func Store(c *fiber.Ctx) error {
	// مرحله ۱: گرفتن ورودی از درخواست
	var user Models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ورودی نامعتبر است",
			"errors":  err.Error(),
		})
	}

	// مرحله ۲: اعتبارسنجی (Validation)
	v := Validations.UserValidation(user)
	if !v.IsValid() {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": "اطلاعات وارد شده معتبر نیست",
			"errors":  v.Errors(),
		})
	}

	// مرحله ۳: ذخیره‌سازی با سرویس
	if err := Services.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "خطا در ذخیره‌سازی کاربر",
			"errors":  err.Error(),
		})
	}

	// مرحله ۴: پاسخ موفقیت
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "کاربر با موفقیت ایجاد شد",
		"data":    user,
	})
}
