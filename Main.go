package main

import (
	"os"

	"github.com/KiyanPourazar/DummyGo/Error"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	Error.CheckError(err)

	appName := os.Getenv("APP_NAME")
	print(appName)

}
