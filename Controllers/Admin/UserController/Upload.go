package UserController

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
)

func UploadProfile(c *fiber.Ctx) error {
	id := c.Params("id")

	// شروع تراکنش
	tx := Config.DB.Begin()
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not start the transaction",
		})
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// پیدا کردن کاربر
	var user Models.User
	if err := tx.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	currentDir, _ := os.Getwd()
	oldFile := filepath.Join(currentDir, user.Image)
	newFile := ""

	// دریافت فایل
	file, err := c.FormFile("image")
	if file != nil && err == nil {
		splittedName := strings.Split(file.Filename, ".")
		extension := splittedName[len(splittedName)-1]
		fileName := fmt.Sprintf("%s.%s", user.Phone, extension)
		relativePath := filepath.Join("uploads", "images", "user-profiles", fileName)
		newFile = filepath.Join(currentDir, relativePath)

		// ساختن پوشه اگر وجود نداشت
		os.MkdirAll(filepath.Dir(newFile), os.ModePerm)

		// ذخیره فایل با متد خود Fiber
		if err := c.SaveFile(file, newFile); err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to save the file",
			})
		}

		// به‌روزرسانی آدرس عکس کاربر
		user.Image = relativePath
	}

	// ذخیره در دیتابیس
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		if newFile != "" {
			os.Remove(newFile)
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "An error occurred while saving data",
		})
	}

	// نهایی کردن تراکنش
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		if newFile != "" {
			os.Remove(newFile)
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "An error occurred while commit data",
		})
	}

	// حذف فایل قدیمی (اگر وجود داشت و جایگزین شد)
	if user.Image != "" && oldFile != newFile {
		if _, err := os.Stat(oldFile); err == nil {
			os.Remove(oldFile)
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "profile updated successfully",
		"image":   user.Image,
	})
}
