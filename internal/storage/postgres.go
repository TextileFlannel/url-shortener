package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) CreateLink(originalURL string) (string, error) {
	var shortURL string
	err := s.db.QueryRow(
		`INSERT INTO links (original_url, short_url) 
		VALUES ($1, $2) 
		ON CONFLICT (original_url) DO NOTHING
		RETURNING short_url`,
		originalURL, GenerateShortURL(),
	).Scan(&shortURL)

	if err != nil {
		if err == sql.ErrNoRows {
			err = s.db.QueryRow(
				`SELECT short_url FROM links WHERE original_url = $1`,
				originalURL,
			).Scan(&shortURL)
			if err != nil {
				return "", fmt.Errorf("failed to get existing short URL: %w", err)
			}
			return shortURL, nil
		}
		return "", fmt.Errorf("failed to create link: %w", err)
	}

	return shortURL, err
}

func (s *PostgresStorage) GetLink(shortURL string) (string, error) {
	var originalURL string
	err := s.db.QueryRow(
		"SELECT original_url FROM links WHERE short_url = $1",
		shortURL,
	).Scan(&originalURL)

	if err == sql.ErrNoRows {
		return "", ErrNotFound
	}
	return originalURL, err
}
