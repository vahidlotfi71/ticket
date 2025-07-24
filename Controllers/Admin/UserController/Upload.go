package UserController

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/ticket/Config"
	"github.com/vahidlotfi71/ticket/Models"
	"github.com/valyala/fasthttp"
)

func UploadProfile(c *fiber.Ctx) error {
	id := c.Params("id")

	//  یک تراکنش دیتابیس ایجاد کردیم
	tx := Config.DB.Begin()

	// اگر وسط کار خطایی پیش آمدهمه تغییرات لغو
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// اگر نتونست تراکنش را شروع کند
	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not start the transaction",
		})
	}

	var user Models.User
	tx.Model(&Models.User{}).First(&user, id)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	current_dir, _ := os.Getwd()
	old_file := filepath.Join(current_dir, user.Image)
	new_file := ""

	file, err := c.FormFile("image")
	if file != nil && err == nil {
		splitted_name := strings.Split(file.Filename, ".")
		extension := splitted_name[len(splitted_name)-1]
		fileName := fmt.Sprintf("%s.%s", user.Phone, extension)
		relative_path := filepath.Join("uploads", "images", "user-profiles", fileName)
		new_file = filepath.Join(current_dir, relative_path)
		if err := fasthttp.SaveMultipartFile(file, new_file); err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to save the file",
			})
		}
		user.Image = relative_path
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		if new_file != "" {
			os.Remove(new_file)
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "An error occurred while saving data",
		})
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		if new_file != "" {
			os.Remove(new_file)
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "An error occurred while commit data",
		})
	}
	_, fileErr := os.Stat(old_file)
	if os.IsNotExist(fileErr) {
		os.Remove(old_file)
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "profile-updated successfully",
	})
}
