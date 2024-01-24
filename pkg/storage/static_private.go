package storage

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/lowl11/boost2/pkg/config"
)

func buildConnectionString(ignoreSchema bool) string {
	host := config.Get("C360_PG_HOST").String()
	port := config.Get("C360_PG_PORT").String()
	databaseName := config.Get("C360_PG_DB").String()
	username := config.Get("C360_PG_USER").String()
	password := config.Get("C360_PG_PASS").String()

	var schemaPart string
	if !ignoreSchema {
		schemaPart = "search_path=" + config.Get("C360_DB_SCHEMA").String()
	}

	cs := strings.Builder{}
	_, _ = fmt.Fprintf(
		&cs,
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable %s",
		host, port, username, password, databaseName, schemaPart,
	)

	pgxCS, _ := pgx.ParseConfig(cs.String())
	return stdlib.RegisterConnConfig(pgxCS)
}
