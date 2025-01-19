package domain

import (
	"errors"
	"time"
)

var (
	ErrMissingParameter = errors.New("missing parameter")
)

type Person struct {
	ID        int       `json:"person_id" db:"id"`
	FirstName string    `json:"firstname" db:"firstname"`
	LastName  string    `json:"lastname" db:"lastname"`
	DNI       int       `json:"dni" db:"dni"`
	BirthDate time.Time `json:"birthdate" db:"birthdate"`
}

// patron Factory: Es una fabrica que rotorna el tipo de entidad
func NewPerson(firstName string, lastName string, DNI int, birthDate time.Time) (*Person, error) {
	if firstName == "" || lastName == "" {
		return nil, ErrMissingParameter
	}
	if birthDate.IsZero() {
		return nil, ErrMissingParameter
	}

	if DNI == 0 {
		return nil, ErrMissingParameter
	}

	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		DNI:       DNI,
		BirthDate: birthDate,
	}, nil
}
