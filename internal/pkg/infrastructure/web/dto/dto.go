package dto

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrMissingParameter    = errors.New("missing parameter")
	ErrInvalidTypeVariable = errors.New("invalid type of variable")
)

type UserCreateDto struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	DNI       int    `json:"dni"`
	BirthDate string `json:"birthdate"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type UserDtoResponse struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	DNI       int    `json:"dni"`
	BirthDate string `json:"birthdate"`
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
} // En este struct deberia ir solo lo que queremos mostrar a la hora de mostrar un user.

func ValidateUserCreateDto(firstName string, lastName string, DNI int, birthDate string, email string, password string, role string) error {

	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" || strings.TrimSpace(role) == "" ||
		strings.TrimSpace(firstName) == "" || strings.TrimSpace(lastName) == "" || strings.TrimSpace(birthDate) == "" || DNI == 0 {
		return ErrMissingParameter
	}
	if reflect.TypeOf(email) != reflect.TypeOf("") || reflect.TypeOf(password) != reflect.TypeOf("") || reflect.TypeOf(role) != reflect.TypeOf("") {
		return ErrInvalidTypeVariable
	}
	return nil
} // Las validadciones estan mal, si mandas " " con un espacio no lo toma como error. Con .TrimSpace() elimino los espacios en blanco iniciales y finales.

func ValidateInputId(id int) error {

	if id == 0 {
		return ErrMissingParameter
	}
	if reflect.TypeOf(id) != reflect.TypeOf(1) {
		return ErrInvalidTypeVariable
	}
	return nil
}
