package configs

import (
	"database/sql"

	"github.com/alifrahmadian/alif-bookcabin-coding-test/internal/handlers"
)

type Config struct {
	DB      *sql.DB
	Handler *Handler
	Env     string
}

type Handler struct {
	SeatMapHandler *handlers.SeatMapHandler
}
