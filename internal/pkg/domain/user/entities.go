package domain

import (
	domain "canchitas-libres-user/internal/pkg/domain/person"
	"errors"
	"strings"
	"time"
)

type User struct {
	Person   *domain.Person //`json:"person" db:"person"`
	Id       int            `json:"id" db:"user_id"`
	Email    string         `json:"email" db:"email"`
	Password string         `json:"password" db:"password"`
	Active   bool           `json:"active" db:"active"`
	Role     string         `json:"role" db:"role"`
}

var (
	ErrMissingParameter      = errors.New("missing parameter")
	ErrRoleInvalid           = errors.New("role doesnt exist")
	ErrPasswordMinCharacters = errors.New("password too short")
	ErrCreatePerson          = errors.New("error to create person")
)

// NewUser creates a new User instance
func NewUser(firstName string, lastName string, DNI int, birthDate time.Time, email string, password string, role string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, ErrMissingParameter
	}
	if birthDate.IsZero() {
		return nil, ErrMissingParameter
	}
	if DNI == 0 {
		return nil, ErrMissingParameter
	} // Estas validaciones ya se hacen en la entidad persona, no deberia hacerlas dos veces.

	person, errPerson := domain.NewPerson(firstName, lastName, DNI, birthDate)
	if errPerson != nil {
		return nil, ErrCreatePerson
	}

	if email == "" || password == "" || role == "" {
		return nil, ErrMissingParameter
	}
	r := strings.ToLower(role)
	if r != "admin" && r != "user" {
		return nil, ErrRoleInvalid
	}
	if len(password) < 5 {
		return nil, ErrPasswordMinCharacters
	}
	return &User{
		Person:   person,
		Email:    email,
		Password: password,
		Active:   true,
		Role:     role,
	}, nil
}

//Defini las variables de la struct publicas porque sino el unmarshall del createUser de handler me decia que no estaba exportando nada.

// Getters
// func (u *User) GetPerson() *domain.Person {
// 	return u.Person
// }

// func (u *User) GetId() string {
// 	return u.id
// }

// func (u *User) GetEmail() string {
// 	return u.email
// }

// func (u *User) GetPassword() string {
// 	return u.password
// }

// func (u *User) IsActive() bool {
// 	return u.active
// }

// func (u *User) GetRole() string {
// 	return u.role
// }

// // Setters
// func (u *User) SetPerson(person *domain.Person) {
// 	u.person = person
// }

// func (u *User) SetId(id string) {
// 	u.id = id
// }

// func (u *User) SetEmail(email string) {
// 	u.email = email
// }

// func (u *User) SetPassword(password string) {
// 	u.password = password
// }

// func (u *User) SetActive(active bool) {
// 	u.active = active
// }

// func (u *User) SetRole(role string) {
// 	u.role = role
// }
