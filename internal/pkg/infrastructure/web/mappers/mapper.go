package mappers

import (
	domain "canchitas-libres-user/internal/pkg/domain/user"
	"canchitas-libres-user/internal/pkg/infrastructure/web/dto"
	"time"
)

func ToDomainUser(dtoUser dto.UserCreateDto) (domain.UserCreateInput, error) {

	layout := "2006-01-02"
	birthDate, _ := time.Parse(layout, dtoUser.BirthDate)

	return domain.UserCreateInput{
		FirstName: dtoUser.FirstName,
		LastName:  dtoUser.LastName,
		DNI:       dtoUser.DNI,
		BirthDate: birthDate,
		Email:     dtoUser.Email,
		Password:  dtoUser.Password,
		Active:    true,
		Role:      dtoUser.Role,
	}, nil
}

func ToDtoUser(domainUser domain.User) (dto.UserDtoResponse, error) {
	return dto.UserDtoResponse{
		FirstName: domainUser.Person.FirstName,
		LastName:  domainUser.Person.LastName,
		DNI:       domainUser.Person.DNI,
		BirthDate: domainUser.Person.BirthDate.Format("2006-01-02"),
		Id:        domainUser.Id,
		Email:     domainUser.Email,
		Role:      domainUser.Role,
	}, nil
}
