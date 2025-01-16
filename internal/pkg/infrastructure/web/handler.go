package web

import (
	"canchitas-libres-field/internal/pkg/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface {
	GetAll() ([]domain.Field, error)
	GetByID(id string) (domain.Field, error)
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

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
        id := r.URL.Query().Get("id")
        if id == "" {  
            handler.GetAllFields(w, r)
            return
        }
        
        handler.GetFieldByID(w, r)
        return
	case r.Method == http.MethodPost:

		handler.CreateField(w, r)
		return
	default:
		http.NotFound(w, r)
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
	id := r.URL.Query().Get("id")
	field, err := handler.Service.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	handler.Service.Delete(field.FieldID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.Write([]byte("Field deleted"))
}
