package storage

import (
	"context"
	"database/sql"
)

type (
	Storage interface {
		CreateUser(ctx context.Context, username string, password string, telegramID string) error
	}

	PostgresStorage struct {
		Db *sql.DB
	}
)
