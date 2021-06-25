package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"yawoen.com/app/internal/config"
)

//#region: configurando o módulo helpers
var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

//#endregion

//#region: errors helpers
type Error struct {
	StatusCode int
	Message    string
}

// Configura um template para o envio dos errors json
func baseError(w http.ResponseWriter, status int, message string) {
	customError := Error{
		StatusCode: status,
		Message:    message,
	}

	jsonError, err := json.MarshalIndent(customError, "", "  ")
	if err != nil {
		app.ErrorLog.Panic(fmt.Sprintf("Some error occurred with marshal %s", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonError)

	var test Error
	json.Unmarshal([]byte(jsonError), &test)
	log.Println(test)
}

// Envia uma mensagem de erro formatada para json
func ClientError(w http.ResponseWriter, status int, message string) {
	app.InfoLog.Println("Client error with status of", status)

	baseError(w, status, message)
}

// Envia uma mensagem de erro formatada para json
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	baseError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

//#endregion

// type jsonResponse struct {
// 	OK      bool   `json:"ok"`
// 	Message string `json:"message"`
// }

// func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
// 	resp := jsonResponse{
// 		OK:      true,
// 		Message: "Available!",
// 	}

// 	out, err := json.MarshalIndent(resp, "", "     ")
// 	if err != nil {
// 		helpers.ServerError(w, err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(out)
// }
