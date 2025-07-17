package routes

import "github.com/gofiber/fiber/v2"

// تعریف کنترلر به‌صورت closure
var defaultController = func(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "hello from default controller",
	})
}

func Routes(app *fiber.App) {
	// گروه اصلی admin
	adminGroup := app.Group("/admin")

	adminGroup.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "admin index route",
		})
	})

	// گروه user زیرمجموعه admin
	userGroup := adminGroup.Group("/user")

	userGroup.Get("/", defaultController)            // لیست همه‌ی کاربران
	userGroup.Get("/show/:id", defaultController)    // نمایش جزئیات کاربر
	userGroup.Post("/store", defaultController)      // ساخت کاربر جدید
	userGroup.Post("/update/:id", defaultController) // ویرایش کاربر
	userGroup.Post("/delete/:id", defaultController) // حذف کاربر
}
