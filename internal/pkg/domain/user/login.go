package domain

import (
	"errors"
)

func (s *Service) Login(email string, password string) (string, error) {

	var ErrUserNotExist = errors.New("invalid credentials")
	var ErrUserBanned = errors.New("user not activated")

	userArray, err := s.StorageRepository.GetAll()
	if err != nil {
		return "", err
	}

	for i := range userArray {
		if userArray[i].Email == email && userArray[i].Password == password {
			if userArray[i].Active {
				return "token: 123", nil //Deberiamos devolver un token JWT. LLamar a una interfaz que nos genere desde la infrastructure el token?
			} else {
				return "", ErrUserBanned
			}
		}
	}
	return "", ErrUserNotExist
}
