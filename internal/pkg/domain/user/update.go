package domain

import (
	"context"
	"errors"
	"strings"
	"time"
)

func (s *Service) Update(id int, userU User) error {
	var errIDNotFound = errors.New("ID not found")

	userU.Person.FirstName = strings.TrimSpace(userU.Person.FirstName)
	userU.Person.LastName = strings.TrimSpace(userU.Person.LastName)
	userU.Email = strings.TrimSpace(userU.Email)
	userU.Password = strings.TrimSpace(userU.Password)
	userU.Role = strings.TrimSpace(userU.Role)

	r := strings.ToLower(userU.Role)
	if r != "admin" && r != "user" && r != "" {
		return ErrRoleInvalid
	}

	if userU.Password != "" && len(userU.Password) < 5 {
		return ErrPasswordMinCharacters
	}

	today := time.Now()
	userYears := today.Year() - userU.Person.BirthDate.Year()

	if today.Month() < userU.Person.BirthDate.Month() || (today.Month() == userU.Person.BirthDate.Month() && today.Day() < userU.Person.BirthDate.Day()) {
		userYears--
	}
	if userU.Person.BirthDate.Format("2006-01-02") != "0001-01-01" && (userYears < 18 || userYears > 130) {
		return ErrMinAge
	} //El birthdate cuando no se le asigna nada, da por defecto 0001-01-01 00:00:00 +0000 UTC. Cuando le das un formato invalido tambien.

	userArray, errGetAll := s.StorageRepository.GetAll()
	if errGetAll != nil {
		return errGetAll
	}

	for i := range userArray {
		if userArray[i].Person.ID == id {
			return s.StorageRepository.Update(context.Background(), id, userU)
		}
	}
	return errIDNotFound
}
