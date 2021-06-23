package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"yawoen.com/app/pkg/configs"
)

// Configura todos os endpoints da aplicação
func setUpRoutes(a *configs.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//#region: middlewares
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	//#endregion

	//#region: route
	// mux.Get("/admin", )

	//#endregion

	//#region: static route
	fileServer := http.FileServer(http.Dir("./public/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	//#endregion

	return mux
}
