package configs

import (
	"fmt"
	"os"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/db"
	"github.com/joho/godotenv"
)

func LoadDBConfig() *db.DBConfig {
	return &db.DBConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func LoadEnv() string {
	return os.Getenv("ENV")
}

func LoadGoDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file")
	}

	return nil
}
