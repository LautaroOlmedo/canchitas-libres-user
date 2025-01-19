package domain

import (
	"canchitas-libres-user/internal/configuration"
	"context"
)

type StorageRepository interface {
	GetAll() ([]User, error) //Desde la base de datos no puede devolver un user porque el user no tiene una columna persona. Devuelvo otra struct creada en el caso de uso?
	GetByID(id int) (User, error)
	Add(ctx context.Context, user User) error
	Delete(ctx context.Context, id int) error
	//Update(ctx context.Context, id string) error
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
