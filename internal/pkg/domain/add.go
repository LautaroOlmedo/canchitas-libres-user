package domain

import (
	"context"
	"fmt"
)

func (s *Service) Add(field Field) error {
	fmt.Println("in use case we have a field whit name:", field.Name)
	return s.StorageRepository.Add(context.Background(), field)
}
