package App

import (
	"os"

	"github.com/KiyanPourazar/DummyGo/Router"
	"github.com/go-chi/chi"
)

var App Application

type Application struct {
	Domain  string
	AppName string
	Router  *chi.Mux
}

func InitApp(app *Application) {
	app.Domain = os.Getenv("APP_DOMAIN")
	app.AppName = os.Getenv("APP_NAME")
	app.Router = Router.Router()
}
