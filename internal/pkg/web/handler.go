package web

import (
	"net/http"
)

type Service interface {
	GetAll() error
	// GetByID
	// Add()
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
	case r.Method == http.MethodPost:
		w.Header().Set("Content-Type", "application/xml; encoding=UTF-8")
		return
	default:
		http.NotFound(w, r)
		return
	}
}

func (handler *Handler) GetAllFields(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (handler *Handler) GetFieldByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (handler *Handler) CreateField(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (handler *Handler) UpdateField(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (handler *Handler) DeleteField(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
