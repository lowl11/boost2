package storage_logger

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost2/internal/infrastructure/config"
)

type Logger struct {
	connection *sqlx.DB
	env        string
	service    string
}

func New(connection *sqlx.DB) *Logger {
	debugMode := config.Get().Get("C360_DEBUG_MODE").Bool()
	var env string
	if debugMode {
		env = "TEST"
	} else {
		env = "PROD"
	}

	return &Logger{
		connection: connection,
		env:        env,
		service:    config.Get().Get("C360_SERVICE_NAME").String(),
	}
}
