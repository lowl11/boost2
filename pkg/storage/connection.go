package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

func ConnectPool(builder *ConnectionBuilder) (*sqlx.DB, error) {
	// open connection pool
	connection, err := sqlx.Open("pgx", buildConnectionString())
	if err != nil {
		return nil, err
	}

	// set important settings
	connection.SetMaxOpenConns(builder.maxOpenConnections)
	connection.SetMaxIdleConns(builder.maxIdleConnections)

	connection.SetConnMaxLifetime(builder.getMaxConnectionLifetime())
	connection.SetConnMaxIdleTime(builder.getMaxIdleLifetime())

	// ping database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err = connection.PingContext(ctx); err != nil {
		return nil, err
	}

	return connection, nil
}
