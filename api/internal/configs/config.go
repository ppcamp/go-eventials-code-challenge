package configs

import (
	"log"
)

type AppConfig struct {
	InfoLog      *log.Logger
	InProduction bool
	// TokenAuth    *jwtauth.JWTAuth
}
