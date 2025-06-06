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
	"strings"
)

type Service interface {
	GetAll() ([]domain.User, error)
	GetByID(id string) (domain.User, error)
	Add(user domain.User) error
	Delete(id string) error
	Update(id string, user domain.User) error
	Login(email string, password string) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
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
	getAllRe  = regexp.MustCompile(`^\/user\/?$`)
	getOneRe  = regexp.MustCompile(`^\/user\/([a-fA-F0-9-]{36})$`)
	createRe  = regexp.MustCompile(`^\/user\/?$`)
	loginPath = regexp.MustCompile(`^\/user\/login$`)
	getByEm   = regexp.MustCompile(`^\/user\/([a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,})$`)
)

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && createRe.MatchString(r.URL.Path):
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
	case r.Method == http.MethodPut:
		handler.UpdateUser(w, r)
		return
	case r.Method == http.MethodPost && loginPath.MatchString(r.URL.Path):
		handler.LoginUser(w, r)
		return
	case r.Method == http.MethodGet && getByEm.MatchString(r.URL.Path):
		handler.GetUserByEmail(w, r)
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

	usersDto := make([]dto.UserDtoResponse, len(users))

	for i := range users {
		usersDto[i], err = mappers.ToDtoUserResponse(users[i])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}

	usersJSON, jsonerr := json.Marshal(usersDto) //Lo transforma en codigo legible para json
	if jsonerr != nil {
		fmt.Println("error en el marshal de user get all")
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
	id := parts[len(parts)-1]                   // Ultimo elemento del slice --> ultimo elemento del path

	err := dto.ValidateInputId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	user, err := handler.Service.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userDto, err := mappers.ToDtoUserResponse(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//Respuesta
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

	userDto := dto.UserDto{}

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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	id := parts[len(parts)-1]

	err := dto.ValidateInputId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	userDto := dto.UserDto{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(body, &userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = dto.ValidateUserCreateDto(userDto.FirstName, userDto.LastName, userDto.DNI, userDto.BirthDate, userDto.Email, userDto.Password, userDto.Role)
	if err == dto.ErrInvalidTypeVariable {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userU, err := mappers.ToDomainUser(userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = handler.Service.Update(id, userU)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//idString := r.URL.Query().Get("id")

	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	id := parts[len(parts)-1]
	// id, err := strconv.Atoi(idString)
	// if err != nil {
	// 	fmt.Println("error al convertir el id en un int")
	// 	return
	// }

	err := dto.ValidateInputId(id)
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
}

func (handler *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dtoLogin := dto.LoginDto{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(body, &dtoLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = dto.ValidateLogin(dtoLogin.Email, dtoLogin.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userLogin, err := handler.Service.Login(dtoLogin.Email, dtoLogin.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} // El caso de uso devuelve directamente el user y desde el handler me ocupo de llamar a la funcion que genera el token JWT.

	// userToken, errToken := authservice.TokenGenerator(userLogin.Id, userLogin.Role) //Directamente llamo a la implementacion del token desde acá.
	// if errToken != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(errToken.Error()))
	// 	return
	// }

	userDto, err := mappers.ToDtoUserResponse(userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

func (handler *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	email := parts[len(parts)-1]

	err := dto.ValidateInputId(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
	}

	user, err := handler.Service.GetByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userDto, err := mappers.ToDtoUserResponse(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
