package web

import (
	"canchitas-libres-field/internal/pkg/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Service interface {
	GetAll() ([]domain.Field, error)
	GetByID(id int) (domain.Field, error)
	Add(field domain.Field) error
	Delete(id string) error
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
	getAllRe = regexp.MustCompile(`^\/[\/]*$`)
	getOneRe = regexp.MustCompile(`^\/(\d+)$`)
)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && getAllRe.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		h.GetAllFields(w, r)
		return
	case r.Method == http.MethodGet && getOneRe.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		h.GetFieldByID(w, r)
		return
	case r.Method == http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		h.CreateField(w, r)
		return
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func (handler *Handler) GetAllFields(w http.ResponseWriter, r *http.Request) {
	fields, err := handler.Service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fieldsJSON, jsonErr := json.Marshal(fields)
	if jsonErr != nil {
		return
	}

	// Configures the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send the JSON response to the client
	_, _ = w.Write(fieldsJSON)

}

func (handler *Handler) GetFieldByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	id := strings.TrimPrefix(r.URL.Path, "/")
	fmt.Println(id)

	w.Write([]byte("Hello World"))
}

func (handler *Handler) CreateField(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	field := domain.Field{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(body, &field)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = handler.Service.Add(field)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("field was created"))
}

func (handler *Handler) UpdateField(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (handler *Handler) DeleteField(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
