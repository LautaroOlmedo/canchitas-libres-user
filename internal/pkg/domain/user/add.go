package domain

import (
	"context"
	"time"
)

type UserCreateInput struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	DNI       int       `json:"dni"`
	BirthDate time.Time `json:"birthdate"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Active    bool      `json:"active"`
	Role      string    `json:"role"`
}

func (s *Service) Add(userInput UserCreateInput) error {
	userNew, err := NewUser(userInput.FirstName, userInput.LastName, userInput.DNI, userInput.BirthDate, userInput.Email, userInput.Password, userInput.Role)
	if err != nil {
		return err
	}
	return s.StorageRepository.Add(context.Background(), *userNew)
}

//Hacer las validaciones aca, no en la entidad
