package mapper

import (
	"canchitas-libres-field/internal/pkg/domain"
	"canchitas-libres-field/internal/pkg/infrastructure/web/dto"
)

func FieldCreateInput(f dto.FieldDTO) domain.FieldCreateInput {
	return domain.FieldCreateInput{
		Name:   f.Name,
		Number: f.Number,
	}
}
