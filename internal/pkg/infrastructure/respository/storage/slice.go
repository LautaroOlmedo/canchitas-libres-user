package storage

import (
	"canchitas-libres-field/internal/pkg/domain"
	"context"
	"fmt"
)

func (s *Slice) GetAll() ([]domain.Field, error) {
	return s.SliceArr, nil
}

func (s *Slice) Add(ctx context.Context, field domain.Field) error {
	fmt.Println("in infrastructure layer we have a field whit name: ", field.Name)
	s.SliceArr = append(s.SliceArr, field)
	return nil
}

func (s *Slice) Delete(ctx context.Context, id string) error {
	return nil
}
