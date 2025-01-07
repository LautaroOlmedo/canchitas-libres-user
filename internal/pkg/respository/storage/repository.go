package storage

import "canchitas-libres-field/internal/configuration"

type Postgres struct {
	Configuration *configuration.Configuration
}
type Slice struct {
	Configuration *configuration.Configuration
}

func NewPostgresStorage(config *configuration.Configuration) *Postgres {
	return &Postgres{
		Configuration: config,
	}
}

func NewSliceStorage(config *configuration.Configuration) *Slice {
	return &Slice{
		Configuration: config,
	}
}
