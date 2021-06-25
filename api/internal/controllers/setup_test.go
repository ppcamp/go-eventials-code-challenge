package controllers

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/subosito/gotenv"
	"yawoen.com/app/internal/config"
	"yawoen.com/app/internal/helpers"
)

var app config.AppConfig

// creates an application
func TestMain(m *testing.M) {
	// Configures the environment values, used to fetch
	gotenv.Load("./../../.env")

	// change this to true when in production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	repo := NewTestRepo(&app)
	helpers.NewHelpers(&app) // registra os helpers na aplicação
	SetRepository(repo)

	os.Exit(m.Run())
}

func setUpRoutes() http.Handler {
	mux := chi.NewRouter()

	//#region middlewares
	// mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	//#endregion

	//#region: endpoints
	mux.Get("/", Repo.Home)

	//#region RESTy routes for company
	mux.Route("/company", func(r chi.Router) {
		r.Post("/", Repo.CompanyPost)             // creates a new element from scratch
		r.Get("/", Repo.CompanyGetMany)           // search by query string
		r.Put("/website", Repo.CompanyPutWebsite) // updates the website based on others params
		// Subrouters:
		r.Route("/{id}", func(subr chi.Router) { //
			subr.Get("/", Repo.CompanyGetOne)    // search for a given id
			subr.Delete("/", Repo.CompanyDelete) // deletes the element that has this id
			subr.Put("/", Repo.CompanyPut)       // edit all the fields to the element that has this id
			// subr.Patch("/", Repo.CompanyDelete)  // TODO: change only the sended elements
		})
	})

	return mux
}
