package storage

import (
	"context"
	"database/sql"
	"errors"
)

func (s *PostgresStorage) CreateUser(ctx context.Context, username string, password string, telegramID string) error {
	_, err := s.Db.QueryContext(ctx, "INSERT INTO users (username, password, telegram_id, balance) VALUES ($1, $2, $3, $4)", username, password, telegramID, 0)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) GetPasswordFromUsername(ctx context.Context, username string) (string, error) {
	var password string
	err := s.Db.QueryRowContext(ctx, "SELECT password FROM users WHERE username = $1", username).Scan(&password)
	if err != nil || password == "" {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNotFound
		}
		return "", err
	}
	return password, nil

}

func (s *PostgresStorage) GetIdFromUsername(ctx context.Context, username string) (int, error) {
	var userID int
	err := s.Db.QueryRowContext(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&userID)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return 0, err
		}
		return 0, ErrNotFound
	}
	return userID, nil
}
