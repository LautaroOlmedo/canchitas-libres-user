package domain

import (
	"context"
)

//simulamos la funcionalidad de la base de datos para que en los unitest de los casos de uso no llamar todo el tiempo a la BD. Tenemos que implementar la interface de storage.
// type StorageRepository interface {
// 	GetAll() ([]User, error)
// 	GetByID(id int) (User, error)
// 	Add(ctx context.Context, user User) error
// 	Delete(ctx context.Context, id int) error
// 	Update(ctx context.Context, id int, userU User) error
// }

type Mock struct{}

func (m Mock) GetAll() ([]User, error) {
	return []User{}, nil
}

func (m Mock) GetByID(id int) (User, error) {
	return User{}, nil
}

func (m Mock) Add(ctx context.Context, user User) error {
	return nil
}

func (m Mock) Delete(ctx context.Context, id int) error {
	return nil
}

func (m Mock) Update(ctx context.Context, id int, userU User) error {
	return nil
}
