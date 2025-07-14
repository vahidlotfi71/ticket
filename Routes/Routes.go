package routes

import "github.com/gofiber/fiber/v3"

func Routes(app *fiber.App) {
	// گروه اصلی admin
	adminGroup := app.Group("/admin")

	adminGroup.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "admin index route",
		})
	})

	// گروه user زیر مجموعه‌ی admin
	userGroup := adminGroup.Group("/user")

	userGroup.Get("/", userIndex)             // لیست همه‌ی کاربران
	userGroup.Get("/show/:id", userShow)      // نمایش جزئیات کاربر
	userGroup.Post("/store", userStore)       // ساخت کاربر جدید
	userGroup.Post("/update/:id", userUpdate) // ویرایش کاربر
	userGroup.Post("/delete/:id", userDelete) // حذف کاربر
}
