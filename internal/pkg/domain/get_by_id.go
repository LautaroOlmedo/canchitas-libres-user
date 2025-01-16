package domain

import "context"

func (s *Service) GetByID(id string) (Field, error) {
	var ctx context.Context
	return s.StorageRepository.GetByID(ctx, id)
}