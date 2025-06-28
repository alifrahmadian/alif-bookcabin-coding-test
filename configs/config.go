package configs

import "database/sql"

type Config struct {
	DB      *sql.DB
	Handler *Handler
	Env     string
}

type Handler struct{}
