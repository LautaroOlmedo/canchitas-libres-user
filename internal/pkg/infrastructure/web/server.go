package web

import (
	"canchitas-libres-user/internal/configuration"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	config  *configuration.Configuration
	handler *Handler
}

func NewServer(configuration *configuration.Configuration, handler *Handler) (*Server, error) {
	return &Server{
		config:  configuration,
		handler: handler,
	}, nil
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/", s.handler.ServeHTTP)
	handlerWithCORS := withCORS(mux)

	fmt.Printf("Server listening on %s%s\n", s.config.SERVER.DOMAIN, s.config.SERVER.SERVER_PORT)
	err := http.ListenAndServe(s.config.SERVER.SERVER_PORT, handlerWithCORS)
	if err != nil {
		log.Fatalf("Failed to start server: %s \n", err)
	}
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Encabezados CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
