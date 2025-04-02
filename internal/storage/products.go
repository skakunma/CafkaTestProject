package storage

import (
	"context"
	"database/sql"
	"errors"
)

func (s *PostgresStorage) CreateProduct(ctx context.Context, name string, price float64, categories []int, descriprion string, author int) error {
	var productID int
	err := s.Db.QueryRowContext(ctx, "INSERT INTO products(title, description, price, author) VALUES ($1, $2, $3, $4) RETURNING id", name, descriprion, price, author).Scan(&productID)
	if err != nil {
		return err
	}
	for _, category := range categories {
		_, err = s.Db.QueryContext(ctx, "INSERT INTO categories_products(id_product, id_categories) VALUES ($1, $2)", productID, category)
		if err != nil {
			if !errors.Is(sql.ErrNoRows, err) {
				return err

			}
			return ErrNotFound
		}
	}
	return nil
}
