package domain

import (
	domain "canchitas-libres-user/internal/pkg/domain/person"
	"errors"
	"time"
	//"database/sql"
)

type User struct {
	Person   *domain.Person //`json:"person" db:"person"`
	Id       int            `json:"id" db:"user_id"`
	Email    string         `json:"email" db:"email"`
	Password string         `json:"password" db:"password"`
	Active   bool           `json:"active" db:"active"`
	Role     string         `json:"role" db:"role"`
	Phone    string         `json:"phone" db:"phone"`
	//Phone    sql.NullString `json:"phone"` Esto para trabajar con string que puedan ser nulos. De momento hice que en la base de datos sea obligatorio el phone
}

var (
	ErrCreatePerson = errors.New("error to create person")
)

// NewUser creates a new User instance
func NewUser(firstName string, lastName string, DNI int, birthDate time.Time, email string, password string, role string, phone string) (*User, error) {
	person, errPerson := domain.NewPerson(firstName, lastName, DNI, birthDate)
	if errPerson != nil {
		return nil, ErrCreatePerson
	}
	return &User{
		Person:   person,
		Email:    email,
		Password: password,
		Active:   true,
		Role:     role,
		Phone:    phone,
	}, nil
}

//Defini las variables de la struct publicas porque sino el unmarshall del createUser de handler me decia que no estaba exportando nada.
