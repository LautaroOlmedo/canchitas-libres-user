package mappers

import (
	domain "canchitas-libres-user/internal/pkg/domain/user"
	"canchitas-libres-user/internal/pkg/infrastructure/web/dto"
	"time"
)

func ToDomainUser(dtoUser dto.UserDto) (domain.User, error) {

	layout := "2006-01-02"
	birthDate, _ := time.Parse(layout, dtoUser.BirthDate)

	user, _ := domain.NewUser(dtoUser.FirstName, dtoUser.LastName, dtoUser.DNI, birthDate, dtoUser.Email, dtoUser.Password, dtoUser.Role, dtoUser.Phone)
	return *user, nil
}

// func ToDtoUser(domainUser domain.User) (dto.UserDtoResponse, error) {
// 	return dto.UserDtoResponse{
// 		FirstName: domainUser.Person.FirstName,
// 		LastName:  domainUser.Person.LastName,
// 		DNI:       domainUser.Person.DNI,
// 		BirthDate: domainUser.Person.BirthDate.Format("2006-01-02"),
// 		Id:        domainUser.Id,
// 		Email:     domainUser.Email,
// 		Role:      domainUser.Role,
// 	}, nil
// } Esto era para mappear de un user domain a un user dto. Esto puede ser util en caso de que tengamos un userDto como response donde no queremos mostrar todos los datos del user.
