package domain

import "errors"

func (s *Service) GetByEmail(email string) (User, error) {
	var ErrEmailNotExist = errors.New("invalid email")
	var ErrUserBanned = errors.New("user not activated")

	userArray, err := s.StorageRepository.GetAll()
	if err != nil {
		return User{}, err
	}

	for i := range userArray {
		if userArray[i].Email == email {
			if userArray[i].Active {
				return userArray[i], nil
			} else {
				return User{}, ErrUserBanned
			}
		}
	}
	return User{}, ErrEmailNotExist
}
