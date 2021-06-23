package configs

import (
	"log"

	"github.com/go-chi/jwtauth/v5"
)

type AppConfig struct {
	InfoLog      *log.Logger
	InProduction bool
	TokenAuth    *jwtauth.JWTAuth
}
