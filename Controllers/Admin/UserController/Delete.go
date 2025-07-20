package UserController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
)

func Delete(c *fiber.Ctx) error {
	// گرفتن id از پارامتر URL
	id := c.Params("id")

	// تعریف متغیر برای ذخیره کاربر پیدا شده
	var user Models.User

	// جستجو در دیتابیس برای یافتن کاربر
	Config.DB.First(&user, id)

	// اگر کاربر پیدا نشد
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "کاربر پیدا نشد",
		})
	}

	// حذف کاربر از دیتابیس
	if err := Config.DB.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "خطا در حذف کاربر",
			"error":   err.Error(),
		})
	}

	// بازگرداندن پیام موفقیت‌آمیز
	return c.Status(200).JSON(fiber.Map{
		"message": "کاربر با موفقیت حذف شد",
	})
}
