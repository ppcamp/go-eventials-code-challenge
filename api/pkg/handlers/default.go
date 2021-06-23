package handlers

import (
	"log"
	"net/http"
)

// Default endpoint
func Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	log.Println("Request incoming from: ", remoteIP)
}
