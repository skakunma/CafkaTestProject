package storage

import (
	"context"
	"database/sql"
)

type (
	Storage interface {
		CreateUser(ctx context.Context, username string, password string, telegramID string) error
		GetPasswordFromUsername(ctx context.Context, username string) (string, error)
		GetIdFromUsername(ctx context.Context, username string) (int, error)
	}

	PostgresStorage struct {
		Db *sql.DB
	}
)
