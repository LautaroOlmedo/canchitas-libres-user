package user

import "canchitas-libres-field/internal/pkg/domain/person"

type User struct {
	ID       uint           `json:"id" db:"id"` // ID compartido con Person
	Person   *person.Person `json:"person" db:"-"`
	Email    string         `json:"email" db:"email"`
	Password string         `json:"password" db:"password"`
	Active   bool           `json:"active" db:"active"`
	Role     string         `json:"role" db:"role"`
}
