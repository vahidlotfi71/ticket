package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Controllers/Admin/UserController"
)

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

	userGroup.Get("/", UserController.Index)                            // لیست همه‌ی کاربران
	userGroup.Get("/show/:id", UserController.Show)                     // نمایش جزئیات کاربر
	userGroup.Post("/store", UserController.Store)                      // ساخت کاربر جدید
	userGroup.Post("/update/:id", UserController.Update)                // ویرایش کاربر
	userGroup.Post("/delete/:id", UserController.Delete)                // حذف کاربر
	userGroup.Post("/upload-profile/:id", UserController.UploadProfile) // آپلود تصویر

	/* Static file rendering */
	app.Static("/", "public")

	/* Not found response */
	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Route Not found",
		})
	})

}
