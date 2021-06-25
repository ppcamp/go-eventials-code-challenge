package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"yawoen.com/app/internal/config"
	"yawoen.com/app/internal/controllers"
)

// Setup endpoints and middlewares
//
// Returns a chi router
func setUpRoutes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//#region middlewares
	// mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	//#endregion

	//#region: endpoints
	mux.Get("/", controllers.Repo.Home)

	//#region RESTy routes for company
	mux.Route("/company", func(r chi.Router) {
		r.Post("/", controllers.Repo.CompanyPost)             // creates a new element from scratch
		r.Get("/", controllers.Repo.CompanyGetMany)           // search by query string
		r.Put("/website", controllers.Repo.CompanyPutWebsite) // updates the website based on others params
		// Subrouters:
		r.Route("/{id}", func(subr chi.Router) { //
			subr.Get("/", controllers.Repo.CompanyGetOne)    // search for a given id
			subr.Delete("/", controllers.Repo.CompanyDelete) // deletes the element that has this id
			subr.Put("/", controllers.Repo.CompanyPut)       // edit all the fields to the element that has this id
			// subr.Patch("/", controllers.Repo.CompanyDelete)  // TODO: change only the sended elements
		})
	})
	//#endregion

	//#endregion

	//#region static endpoint
	// fileServer := http.FileServer(http.Dir("./api/public/"))
	// mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	//#endregion

	return mux
}
