package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KiyanPourazar/DummyGo/App"
	"github.com/KiyanPourazar/DummyGo/Error"
	"github.com/KiyanPourazar/DummyGo/Router"
	"github.com/joho/godotenv"
)

func main() {
	App.InitApp(&App.App)

	err := godotenv.Load()
	Error.CheckError(err)

	appName := os.Getenv("APP_NAME")
	print(appName)

	log.Printf("The app start in port %v", os.Getenv("APP_PORT"))
	err = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("APP_PORT")), App.App.Router)
	Error.CheckError(err)

	router := Router.Router()

	v1router := Router.Router()
	v1router.HandleFunc("/healthz", handlerReadiness)
	router.Mount("/v1", v1router)

}
