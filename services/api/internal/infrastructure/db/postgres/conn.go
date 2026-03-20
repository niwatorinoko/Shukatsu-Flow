package postgres

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq"

	"shukatsu-flow/api/internal/config"
)

var ErrDatabaseURLIsRequired = errors.New("DATABASE_URL is required")

func NewConnection() (*sql.DB, error) {
	databaseURL := config.GetEnv("DATABASE_URL", "")
	if databaseURL == "" {
		return nil, ErrDatabaseURLIsRequired
	}

	databaseConnection, openError := sql.Open("postgres", databaseURL)
	if openError != nil {
		return nil, openError
	}

	configureConnectionPool(databaseConnection)

	pingError := databaseConnection.Ping()
	if pingError != nil {
		closeError := databaseConnection.Close()
		if closeError != nil {
			return nil, closeError
		}
		return nil, pingError
	}

	return databaseConnection, nil
}

func configureConnectionPool(databaseConnection *sql.DB) {
	databaseConnection.SetMaxOpenConns(10)
	databaseConnection.SetMaxIdleConns(5)
	databaseConnection.SetConnMaxLifetime(5 * time.Minute)
	databaseConnection.SetConnMaxIdleTime(2 * time.Minute)
}
