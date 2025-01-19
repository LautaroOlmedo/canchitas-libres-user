package domain

import (
	"errors"
)

func (s *Service) GetByID(id int) (User, error) {

	var errIDNotFound = errors.New("ID not found")

	userArray, errGetAll := s.StorageRepository.GetAll()
	if errGetAll != nil {
		return User{}, errGetAll
	}

	for i := range userArray {
		if userArray[i].Person.ID == id {
			return s.StorageRepository.GetByID(id)
		}
	}
	return User{}, errIDNotFound //Aca esta bien devolver un user vacio?

}
