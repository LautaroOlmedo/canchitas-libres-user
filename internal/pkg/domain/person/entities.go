package person

import "time"

type Person struct {
	ID        uint      `json:"id" db:"id"`
	FirstName string    `json:"firstname" db:"firstname"`
	LastName  string    `json:"lastname" db:"lastname"`
	DNI       int       `json:"dni" db:"dni"`
	BirthDate time.Time `json:"birthdate" db:"birthdate"`
}
