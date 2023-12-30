package main

import (
	"encoding/json"
	"net/http"

	"github.com/KiyanPourazar/DummyGo/Error"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	Error.CheckError(err)

	w.Header().Add("Conternt-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
