package storage

import (
	domain "canchitas-libres-user/internal/pkg/domain/user"
	"context"
	"fmt"
)

func (s *Slice) GetAll() ([]domain.User, error) {
	return s.SliceArr, nil
}

func (s *Slice) Delete(ctx context.Context, id int) error {
	for i := range s.SliceArr {
		if id == s.SliceArr[i].Person.ID {
			s.SliceArr = append(s.SliceArr[:i], s.SliceArr[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("element with ID %d not found", id)
}

func (s *Slice) Add(ctx context.Context, user domain.User) error {
	//fmt.Println("in infrastructure layer we have a person whit name: ", user.Person.FirstName)
	s.SliceArr = append(s.SliceArr, user)
	return nil
}

func (s *Slice) GetByID(id int) (domain.User, error) {
	for i := range s.SliceArr {
		if id == s.SliceArr[i].Person.ID {
			return s.SliceArr[i], nil
		}
	}
	return domain.User{}, fmt.Errorf("element with ID %d not found", id)
}
