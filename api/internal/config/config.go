package config

import (
	"log"
)

// Holds the app default type wich will be used in the entire application
// This approach prevents circular deps on golang, which isn't supported
type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	// TokenAuth    *jwtauth.JWTAuth
}
