package service

import (
	"context"

	"url-shortener/api"
	"url-shortener/internal/storage"
)

type Service struct {
	api.UnimplementedLinkServiceServer
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) CreateLink(ctx context.Context, req *api.CreateLinkRequest) (*api.CreateLinkResponse, error) {
	short, err := s.storage.CreateLink(req.OriginalUrl)
	if err != nil {
		return nil, err
	}
	return &api.CreateLinkResponse{ShortUrl: short}, nil
}

func (s *Service) GetLink(ctx context.Context, req *api.GetLinkRequest) (*api.GetLinkResponse, error) {
	original, err := s.storage.GetLink(req.ShortUrl)
	if err != nil {
		return nil, err
	}
	return &api.GetLinkResponse{OriginalUrl: original}, nil
}
