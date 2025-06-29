package storage

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestPostgresStorage(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	storage := NewPostgresStorage(db)

	originalUrl := "https://example.com"
	shortUrl := "abc123"

	t.Run("create link successfully", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(
			`INSERT INTO links (original_url, short_url) 
		VALUES ($1, $2) 
		ON CONFLICT (original_url) DO NOTHING
		RETURNING short_url`,
		)).WithArgs(originalUrl, sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"shortUrl"}).AddRow(shortUrl))

		res, err := storage.CreateLink(originalUrl)
		assert.NoError(t, err)
		assert.Equal(t, shortUrl, res)
	})

	t.Run("get link successfully", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(
			"SELECT original_url FROM links WHERE short_url = $1",
		)).WithArgs(shortUrl).WillReturnRows(sqlmock.NewRows([]string{"original_url"}).AddRow(originalUrl))

		res, err := storage.GetLink(shortUrl)
		assert.NoError(t, err)
		assert.Equal(t, originalUrl, res)
	})

	t.Run("get link not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(
			"SELECT original_url FROM links WHERE short_url = $1",
		)).WithArgs(shortUrl).WillReturnError(sql.ErrNoRows)

		res, err := storage.GetLink(shortUrl)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
