package storage

import (
	"canchitas-libres-field/internal/configuration"
	"canchitas-libres-field/internal/pkg/domain"
)

type Postgres struct {
	Configuration *configuration.Configuration
}
type Slice struct {
	Configuration *configuration.Configuration
	SliceArr      []domain.Field
}

func NewPostgresStorage(config *configuration.Configuration) *Postgres {
	return &Postgres{
		Configuration: config,
	}
}

func NewSliceStorage(config *configuration.Configuration) *Slice {
	return &Slice{
		Configuration: config,
		SliceArr:      make([]domain.Field, 0),
	}
}
