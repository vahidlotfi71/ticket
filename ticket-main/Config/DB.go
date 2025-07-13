package Config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
