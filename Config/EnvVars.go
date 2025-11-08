package Config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	JWT_KEY     string
)

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

func GetEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	JWT_KEY = os.Getenv("JWT_KEY")
	return nil
}
