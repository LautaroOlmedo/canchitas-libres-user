package storage

import (
	"canchitas-libres-user/internal/configuration"
	domain "canchitas-libres-user/internal/pkg/domain/user"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	Configuration *configuration.Configuration
	*sqlx.DB
}
type Slice struct {
	Configuration *configuration.Configuration
	SliceArr      []domain.User
}

func NewPostgresStorage(config *configuration.Configuration, db *sqlx.DB) *Postgres {
	return &Postgres{
		Configuration: config,
		DB:            db,
	}
}

func NewSliceStorage(config *configuration.Configuration) *Slice {
	return &Slice{
		Configuration: config,
	}
}
