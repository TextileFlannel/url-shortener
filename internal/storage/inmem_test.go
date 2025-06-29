package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemStorage_CreateLink(t *testing.T) {
	storage := NewInMemStorage()
	origUrl := "http://example.com"

	shortUrl, err := storage.CreateLink(origUrl)
	assert.NoError(t, err)
	assert.NotEmpty(t, shortUrl, "short url should not be empty")

	shortUrl2, err := storage.CreateLink(origUrl)
	assert.NoError(t, err)
	assert.NotEmpty(t, shortUrl2, "short url should not be empty")
}

func TestInMemStorage_GetLink(t *testing.T) {
	storage := NewInMemStorage()
	origUrl := "http://example.com"
	shortUrl, err := storage.CreateLink(origUrl)
	assert.NoError(t, err)

	testUrl, err := storage.GetLink(shortUrl)
	assert.NoError(t, err)
	assert.Equal(t, origUrl, testUrl)

	_, err = storage.GetLink("shortUrl")
	assert.Error(t, err)
}
