package main

import (
	"fmt"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/configs"
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/db"
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/handlers"
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/repositories"
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/routes"
	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/services"
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

	cabinRepo := repositories.NewCabinRepository(db)
	frequentFlyerRepo := repositories.NewFrequentFlyerRepository(db)
	passengerRepo := repositories.NewPassengerRepository(db)
	passengerSeatMapRepo := repositories.NewPassengerSeatMapRepository(db)
	seatMapRepo := repositories.NewSeatMapRepository(db)
	seatRepo := repositories.NewSeatRepository(db)
	seatRowRepo := repositories.NewSeatRowRepository(db)
	seatsItineraryPartRepo := repositories.NewSeatsItineraryPartRepository(db)
	segmentRepo := repositories.NewSegmentRepository(db)
	segmentSeatMapRepo := repositories.NewSegmentSeatMapRepository(db)

	seatMapService := services.NewSeatMapService(
		cabinRepo,
		frequentFlyerRepo,
		passengerRepo,
		passengerSeatMapRepo,
		seatMapRepo,
		seatRepo,
		seatRowRepo,
		seatsItineraryPartRepo,
		segmentRepo,
		segmentSeatMapRepo,
	)

	seatMapHandler := handlers.NewSeatMapHandler(&seatMapService)

	return &configs.Config{
		DB:  db,
		Env: env,
		Handler: &configs.Handler{
			SeatMapHandler: seatMapHandler,
		},
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
