package domain

import (
	"canchitas-libres-user/internal/configuration"
	"context"
)

//go:generate mockery --name=StorageRepository --output=user --inpackage=true
type StorageRepository interface {
	GetAll() ([]User, error)
	GetByID(id string) (User, error)
	Add(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, userU User) error
}

// type Authentication interface {
// 	TokenGenerator(id int, role string) (string, error)
// }

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
