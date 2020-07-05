package transport

import (
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

// TODO, 24.05.2020: move to own package along with whole http initialization
func defaultCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"}, // TODO: Select concrete headers
	})
}

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(defaultCors().Handler)

	return router
}