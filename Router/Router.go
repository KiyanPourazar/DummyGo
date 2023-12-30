package Router

import (
	"github.com/KiyanPourazar/DummyGo/Middleware"
	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	router := chi.NewRouter()
	Middleware.CorsMiddleware(router)

	return router
}
