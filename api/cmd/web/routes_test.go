package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"yawoen.com/app/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := setUpRoutes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}
