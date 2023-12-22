package storage

import (
	"fmt"
	"github.com/lowl11/boost2/pkg/config"
	"strings"
)

func buildConnectionString() string {
	host := config.Get("C360_PG_HOST").String()
	port := config.Get("C360_PG_PORT").String()
	databaseName := config.Get("C360_PG_DB").String()
	username := config.Get("C360_PG_USER").String()
	password := config.Get("C360_PG_PASS").String()
	schema := config.Get("C360_DB_SCHEMA").String()

	cs := strings.Builder{}
	_, _ = fmt.Fprintf(
		&cs,
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, username, password, databaseName, schema,
	)

	return cs.String()
}
