package storage

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("username is not found")

func NewStorage(dns string) (Storage, error) {
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}

	storage := &PostgresStorage{Db: db}
	err = storage.createSchema()

	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *PostgresStorage) createSchema() error {

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username varchar(255) UNIQUE,
		password varchar(255),
	    telegram_id varchar(255) UNIQUE
	    balance float,
	);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    title varchar(255),
    description varchar(255),
    price float
    author int
    FOREIGN KEY (author) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    title varchar(255)
);

CREATE TABLE IF NOT EXISTS categories_products(
    id_product int,
    id_categories int,
    FOREIGN KEY (id_product) REFERENCES products(id) ON DELETE CASCADE
    FOREIGN KEY (id_categories) REFERENCES categories(id) ON DELETE CASCADE
);
`
	_, err := s.Db.Query(query)
	if err != nil {
		return err
	}
	return nil
}
