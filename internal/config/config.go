package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port           string
	DBName         string
	DBUsername     string
	DBPassword     string
	DBRootPassword string
}

func NewConfig() *Config {
	godotenv.Load()
	return &Config{
		Port:           os.Getenv("PORT"),
		DBName:         os.Getenv("DB_NAME"),
		DBUsername:     os.Getenv("DB_USERNAME"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBRootPassword: os.Getenv("DB_ROOT_PASSWORD"),
	}
}
