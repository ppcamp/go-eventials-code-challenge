package config

import (
	"log"
)

type AppConfig struct {
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	// TokenAuth    *jwtauth.JWTAuth
}
