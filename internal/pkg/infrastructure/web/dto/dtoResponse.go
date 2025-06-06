package dto

import (
	"time"
)

type UserDtoResponse struct {
	Id        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	DNI       int       `json:"DNI"`
	Birthdate time.Time `json:"birthday"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Active    bool      `json:"active"`
	Role      string    `json:"role"`
	Phone     string    `json:"phone"`
}
