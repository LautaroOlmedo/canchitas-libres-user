package storage

import (
	"canchitas-libres-field/internal/configuration"
	"github.com/jmoiron/sqlx"
)

type Slice struct {
	Configuration *configuration.Configuration
}

func NewSliceStorage(config *configuration.Configuration) *Slice {
	return &Slice{
		Configuration: config,
	}
}

type Postgres struct {
	Configuration *configuration.Configuration
	*sqlx.DB
}

func NewPostgresStorage(config *configuration.Configuration, db *sqlx.DB) *Postgres {
	return &Postgres{
		Configuration: config,
		DB:            db,
	}
}
