package domain

import (
	"context"
)

func (s *Service) Add(field Field) error {

	return s.StorageRepository.Add(context.Background(), field)
}
