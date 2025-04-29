package domain

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

// type UserInput struct {
// 	FirstName string    `json:"firstname"`
// 	LastName  string    `json:"lastname"`
// 	DNI       int       `json:"dni"`
// 	BirthDate time.Time `json:"birthdate"`
// 	Email     string    `json:"email"`
// 	Password  string    `json:"password"`
// 	Role      string    `json:"role"`
// } //Esta struct es necesaria? Creo que la hice para no crear un user sin validar todo, pero no hay problema en crearlo mientras no lo mandemos a la base de datos.

var (
	ErrMissingParameter      = errors.New("missing parameter")
	ErrRoleInvalid           = errors.New("role doesnt exist")
	ErrPasswordMinCharacters = errors.New("password too short")
	ErrMinAge                = errors.New("age not allow")
)

func (s *Service) Add(user User) error {

	r := strings.ToLower(user.Role)
	if r != "admin" && r != "user" {
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

	return s.StorageRepository.Add(context.Background(), user)
}
