package configs

import (
	"encoding/json"
	"log"
	"net/http"
)

// Uso simplesmente para retornar um json informando o código e a mensagem
// Um atalho apenas
type Error struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e *Error) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.StatusCode)

	jsonfile, err := json.MarshalIndent(e, "", "   ")
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonfile)
}
