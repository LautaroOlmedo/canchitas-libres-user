package domain

import (
	"context"
	"fmt"
)

func (s *Service) Delete(id int) error {
	userArray, err := s.StorageRepository.GetAll()
	if err != nil {
		return err
	}
	for i := range userArray {
		if userArray[i].Person.ID == id {
			return s.StorageRepository.Delete(context.Background(), id)
		}
	}
	return fmt.Errorf("element with ID %d not found", id)
}
