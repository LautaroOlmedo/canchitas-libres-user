package domain

import (
	"context"
	"errors"
)

type FieldCreateInput struct {
	Name   string `query:"name"`
	Number int    `query:"number"`
}

func (s *Service) Add(input FieldCreateInput) error {
	if input.Name == "lautaro" {
		return errors.New("lautaro is not yet supported")
	}
	if input.Number < 1 {
		return errors.New("Number must be greater than zero")
	}
	field := NewField(input.Name, input.Number)
	return s.StorageRepository.Add(context.Background(), field)
}
