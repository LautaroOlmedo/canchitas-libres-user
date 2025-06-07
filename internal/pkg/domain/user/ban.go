package domain

import "context"

func (s *Service) Ban(id string) error {
	user, err := s.GetByID(id)
	if err != nil {
		return err
	}

	if user.Active {
		user.Active = false
	} else {
		user.Active = true
	}

	return s.StorageRepository.UpdateActive(context.Background(), id, user.Active)

}
