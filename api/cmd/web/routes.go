package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"yawoen.com/app/internal/config"
	"yawoen.com/app/internal/controllers"
)

// Configura: todos os endpoints da aplicação;
// e todos os middlewares utilizados
func setUpRoutes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//#region middlewares
	// mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	//#endregion

	//#region: routes
	// rotas globais
	mux.Get("/", controllers.Repo.Home)
	//#region RESTy routes for company
	mux.Route("/company", func(r chi.Router) {
		r.Post("/", controllers.Repo.CompaniesPost)   // cria um elemento
		r.Get("/", controllers.Repo.CompaniesGetMany) // busca por query os elementos
		// Subrouters:
		r.Route("/{id}", func(subr chi.Router) {
			subr.Get("/", controllers.Repo.CompanyGetOne)      // busca por id
			subr.Patch("/", controllers.Repo.CompaniesDelete)  // altera elementos enviados na query
			subr.Put("/", controllers.Repo.CompaniesPut)       // altera todo elemento
			subr.Delete("/", controllers.Repo.CompaniesDelete) // deleta o elemento que tem este id
		})
	})
	//#endregion

	//#endregion

	//#region static route
	// fileServer := http.FileServer(http.Dir("./api/public/"))
	// mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	//#endregion

	return mux
}
