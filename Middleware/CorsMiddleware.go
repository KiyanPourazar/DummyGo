package Middleware

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func CorsMiddleware(router *chi.Mux) {
	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	))
}
