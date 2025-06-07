package domain

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrMissingParameter      = errors.New("missing parameter")
	ErrRoleInvalid           = errors.New("role doesnt exist")
	ErrPasswordMinCharacters = errors.New("password too short")
	ErrMinAge                = errors.New("age not allow")
)

func (s *Service) Add(user User) error {

	user.Role = strings.ToLower(user.Role)
	if user.Role != "admin" && user.Role != "user" && user.Role != "" {
		return ErrRoleInvalid
	}

	if len(user.Password) < 5 {
		return ErrPasswordMinCharacters
	}

	today := time.Now()
	userYears := today.Year() - user.Person.BirthDate.Year()
	if today.Month() < user.Person.BirthDate.Month() || (today.Month() == user.Person.BirthDate.Month() && today.Day() < user.Person.BirthDate.Day()) {
		userYears--
	}
	fmt.Println(userYears)
	if userYears < 18 || userYears > 130 {
		return ErrMinAge
	}

	if user.Role == "" {
		user.Role = "user"
	}

	return s.StorageRepository.Add(context.Background(), user)
}
