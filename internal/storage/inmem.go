package storage

import (
	"sync"
)

type InMemStorage struct {
	mu      sync.RWMutex
	links   map[string]string
	reverse map[string]string
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		links:   make(map[string]string),
		reverse: make(map[string]string),
	}
}

func (s *InMemStorage) CreateLink(originalURL string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if short, exists := s.reverse[originalURL]; exists {
		return short, nil
	}

	short := GenerateShortURL()
	s.links[short] = originalURL
	s.reverse[originalURL] = short
	return short, nil
}

func (s *InMemStorage) GetLink(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	original, exists := s.links[shortURL]
	if !exists {
		return "", ErrNotFound
	}
	return original, nil
}
