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

type UserDto struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	DNI       int    `json:"dni"`
	BirthDate string `json:"birthdate"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Phone     string `json:"phone"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidateUserCreateDto(firstName string, lastName string, DNI int, birthDate string, email string, password string, role string) error {

	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" ||
		strings.TrimSpace(firstName) == "" || strings.TrimSpace(lastName) == "" || strings.TrimSpace(birthDate) == "" || DNI == 0 {
		return ErrMissingParameter
	}
	if reflect.TypeOf(email) != reflect.TypeOf("") || reflect.TypeOf(password) != reflect.TypeOf("") || reflect.TypeOf(role) != reflect.TypeOf("") ||
		reflect.TypeOf(firstName) != reflect.TypeOf("") || reflect.TypeOf(lastName) != reflect.TypeOf("") || reflect.TypeOf(DNI) != reflect.TypeOf(1) ||
		reflect.TypeOf(birthDate) != reflect.TypeOf("") {
		return ErrInvalidTypeVariable
	}
	return nil
} //Con .TrimSpace() elimino los espacios en blanco iniciales y finales.
// No hice ninguna validacion con el telefono porq no es obligatorio ponerlo. Pero algo se deberia hacer.

func ValidateInputId(id string) error {

	if strings.TrimSpace(id) == "" {
		return ErrMissingParameter
	}
	if reflect.TypeOf(id) != reflect.TypeOf("") {
		return ErrInvalidTypeVariable
	}
	return nil
}

func ValidateLogin(email string, password string) error {
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return ErrMissingParameter
	}

	if reflect.TypeOf(email) != reflect.TypeOf("") || reflect.TypeOf(password) != reflect.TypeOf("") {
		return ErrInvalidTypeVariable
	}
	return nil
}
