package storage

import (
	"context"
	"database/sql"
)

type (
	Storage interface {
		//Users
		CreateUser(ctx context.Context, username string, password string, telegramID string) error
		GetPasswordFromUsername(ctx context.Context, username string) (string, error)
		GetIdFromUsername(ctx context.Context, username string) (int, error)
		//Products
		CreateProduct(ctx context.Context, name string, price float64, categories []int, description string) error
	}

	PostgresStorage struct {
		Db *sql.DB
	}
)
