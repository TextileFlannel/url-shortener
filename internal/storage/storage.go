package storage

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type Storage interface {
	CreateLink(originalURL string) (string, error)
	GetLink(shortURL string) (string, error)
}
