package Config

import "os"

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_DATABASE string
	DB_USERNAME string
	DB_PASSWORD string
}

func LoadEnvConfig() Config {
	return Config{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}
}
