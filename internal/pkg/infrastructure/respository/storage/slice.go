package storage

import (
	"canchitas-libres-field/internal/pkg/domain"
	"context"
	"fmt"
)

func (s *Slice) GetAll() ([]domain.Field, error) {
	return s.SliceArr, nil
}

func (s *Slice) Add(ctx context.Context, field domain.Field) error {
	fmt.Println("in infrastructure layer we have a field whit name: ", field.Name)
	s.SliceArr = append(s.SliceArr, field)
	return nil
}

func (s *Slice) Delete(ctx context.Context, id string) error {
	
	for i := range s.SliceArr {
		
		if  s.SliceArr[i].FieldID == id {
			s.SliceArr = append(s.SliceArr[:i], s.SliceArr[i+1:]...)
			return nil
		}
	}
	
	
	return nil
}

func (s *Slice) GetByID(ctx context.Context, id string) (domain.Field , error) {
	for i , _ := range s.SliceArr{
		if s.SliceArr[i].FieldID == id {
			return s.SliceArr[i], nil
		}
	}
	return domain.Field{} , nil 
	
}