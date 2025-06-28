package main

import (
	"fmt"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/configs"
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/db"
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/routes"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config *configs.Config
}

func LoadConfig() (*configs.Config, error) {
	err := configs.LoadGoDotEnv()
	if err != nil {
		return nil, err
	}

	dbConfig := configs.LoadDBConfig()
	env := configs.LoadEnv()

	db, err := db.Connect(*dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return &configs.Config{
		DB:      db,
		Env:     env,
		Handler: &configs.Handler{},
	}, nil
}

func NewApp() *App {
	cfg, err := LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	router := gin.Default()
	routes.SetupRoutes(router, cfg.Handler)

	return &App{
		Router: router,
		Config: cfg,
	}
}
