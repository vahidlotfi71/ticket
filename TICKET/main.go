package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/vahidlotfi71/ticket/Config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// app := fiber.New()

	// app.Get("/Hello", func(c fiber.Ctx) error {
	// 	return c.SendString("hello world ")
	// })
	// app.Listen(":8080")

	err := godotenv.Load("example.env")
	if err != nil {
		log.Fatal("can not load env file : ", err)
	}

	config := Config.LoadEnvConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_DATABASE,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(" Failed to connect to the database:  ", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	fmt.Print("Table created successfully")

}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}
