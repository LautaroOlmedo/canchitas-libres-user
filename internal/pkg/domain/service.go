package domain

import (
	"canchitas-libres-field/internal/configuration"
	"context"
)

type StorageRepository interface {
	GetAll() error
	// GetByID
	// Add()
	Delete(ctx context.Context, id string) error
	// Update
}

type Service struct {
	Config            *configuration.Configuration
	StorageRepository StorageRepository
}

func NewService(config *configuration.Configuration, storageRepository StorageRepository) *Service {
	return &Service{
		Config:            config,
		StorageRepository: storageRepository,
	}
}
