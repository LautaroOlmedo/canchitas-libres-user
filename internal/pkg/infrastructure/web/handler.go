package web

import (
	domain "canchitas-libres-user/internal/pkg/domain/user"
	"canchitas-libres-user/internal/pkg/infrastructure/web/dto"
	"canchitas-libres-user/internal/pkg/infrastructure/web/mappers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Service interface {
	GetAll() ([]domain.User, error)
	GetByID(id int) (domain.User, error)
	Add(user domain.UserCreateInput) error
	Delete(id int) error
	// Update
}
type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

var (
	// getAllRe = regexp.MustCompile(`^\/[\/]*$`)
	// getOneRe = regexp.MustCompile(`^\/(\d+)$`)
	getAllRe = regexp.MustCompile(`^\/user\/?$`)
	getOneRe = regexp.MustCompile(`^\/user\/(\d+)$`)
)

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:
		handler.CreateUser(w, r)
		return
	case r.Method == http.MethodGet && getAllRe.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		handler.GetAllUser(w, r)
		return
	case r.Method == http.MethodGet && getOneRe.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		handler.GetUserByID(w, r)
		return
	case r.Method == http.MethodDelete:
		handler.DeleteUser(w, r)
		return
	default:
		http.NotFound(w, r)
		return
	}
}

func (handler *Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := handler.Service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	usersDto := make([]dto.UserDtoResponse, len(users)) //Transformo el slice de users en uno de dtoUsers
	for i := range users {
		usersDto[i], err = mappers.ToDtoUser(users[i])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	usersJSON, jsonerr := json.Marshal(usersDto) //Lo transforma en codigo legible para json
	if jsonerr != nil {
		return //retornar un error
	}

	w.Header().Set("Content-Type", "application/json") //Avisamos que vamos a trabajar con JSON
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(usersJSON) // enviamos la respuesta en json al cliente
}

func (handler *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/") // Toma todo lo que este en el path despues del primer / --> user/:id
	parts := strings.Split(path, "/")           // Arma un slice con los elementos utilizando / como separador --> ["user", "9858"]
	idString := parts[len(parts)-1]             // Ultimo elemento del slice --> ultimo elemento del path
	id, err := strconv.Atoi(idString)           // Convierto el string en un int
	if err != nil {
		fmt.Println("error al convertir el id en un int")
		return
	}

	err = dto.ValidateInputId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	user, err := handler.Service.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userDto, err := mappers.ToDtoUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	userJson, errJson := json.Marshal(userDto)
	if errJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errJson.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(userJson)

}

func (handler *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userDto := dto.UserCreateDto{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(body, &userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	fmt.Println(userDto)

	err = dto.ValidateUserCreateDto(userDto.FirstName, userDto.LastName, userDto.DNI, userDto.BirthDate, userDto.Email, userDto.Password, userDto.Role)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	userDomain, err := mappers.ToDomainUser(userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = handler.Service.Add(userDomain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user was created"))
}

func (handler *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (handler *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//idString := r.URL.Query().Get("id")

	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	idString := parts[len(parts)-1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("error al convertir el id en un int")
		return
	}

	err = dto.ValidateInputId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	err = handler.Service.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user was eliminated"))
}
