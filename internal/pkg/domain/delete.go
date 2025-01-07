package domain

import "context"

func (s *Service) Delete(id string) error {
	var ctx context.Context
	err := s.StorageRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
