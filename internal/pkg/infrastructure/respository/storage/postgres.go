package storage

import (
	"canchitas-libres-field/internal/pkg/domain"
	"context"
	"fmt"
)

func (s *Postgres) GetAll() ([]domain.Field, error) {
	return []domain.Field{}, nil
}

func (s *Postgres) Add(ctx context.Context, field domain.Field) error {
	fmt.Println("in infrastructure layer we have a field whit name: ", field.Name)

	return nil
}

func (s *Postgres) GetByID(id int) (domain.Field, error) {
	return domain.Field{}, nil
}

func (s *Postgres) Delete(ctx context.Context, id string) error {
	return nil
}
