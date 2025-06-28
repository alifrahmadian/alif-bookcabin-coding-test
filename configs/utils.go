package configs

import (
	"os"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/db"
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
