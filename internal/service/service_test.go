package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"url-shortener/api"
)

type MockStorage struct {
	CreateLinkFunc func(string) (string, error)
	GetLinkFunc    func(string) (string, error)
}

func (m *MockStorage) CreateLink(originalURL string) (string, error) {
	return m.CreateLinkFunc(originalURL)
}

func (m *MockStorage) GetLink(shortURL string) (string, error) {
	return m.GetLinkFunc(shortURL)
}

func TestService_CreateLink(t *testing.T) {
	cntl := gomock.NewController(t)
	defer cntl.Finish()

	t.Run("success creation", func(t *testing.T) {
		mockStorage := &MockStorage{
			CreateLinkFunc: func(originalURL string) (string, error) {
				return "abc123", nil
			},
		}
		srv := NewService(mockStorage)
		resp, err := srv.CreateLink(context.Background(), &api.CreateLinkRequest{
			OriginalUrl: "https://example.com",
		})
		assert.NoError(t, err)
		assert.Equal(t, "abc123", resp.ShortUrl)
	})

	t.Run("storage error", func(t *testing.T) {
		mockStorage := &MockStorage{
			CreateLinkFunc: func(originalURL string) (string, error) {
				return "", assert.AnError
			},
		}
		srv := NewService(mockStorage)
		resp, err := srv.CreateLink(context.Background(), &api.CreateLinkRequest{
			OriginalUrl: "https://example.com",
		})

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestService_GetLink(t *testing.T) {
	cntl := gomock.NewController(t)
	defer cntl.Finish()

	t.Run("success get", func(t *testing.T) {
		mockStorage := &MockStorage{
			GetLinkFunc: func(shortURL string) (string, error) {
				return "https://example.com", nil
			},
		}
		srv := NewService(mockStorage)
		resp, err := srv.GetLink(context.Background(), &api.GetLinkRequest{
			ShortUrl: "abc123",
		})

		assert.NoError(t, err)
		assert.Equal(t, "https://example.com", resp.OriginalUrl)
	})

	t.Run("not found", func(t *testing.T) {
		mockStorage := &MockStorage{
			GetLinkFunc: func(shortURL string) (string, error) {
				return "", errors.New("not found")
			},
		}
		srv := NewService(mockStorage)
		resp, err := srv.GetLink(context.Background(), &api.GetLinkRequest{
			ShortUrl: "abc123",
		})
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("storage error", func(t *testing.T) {
		mockStorage := &MockStorage{
			GetLinkFunc: func(shortURL string) (string, error) {
				return "", assert.AnError
			},
		}
		srv := NewService(mockStorage)
		resp, err := srv.GetLink(context.Background(), &api.GetLinkRequest{
			ShortUrl: "abc123",
		})

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
