package controllers

import (
	"net/http"

	"yawoen.com/app/internal/config"
	"yawoen.com/app/internal/driver"
	"yawoen.com/app/internal/repository"
	"yawoen.com/app/internal/repository/dbrepo"
)

//#region: Repository config

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepository
}

// Repo the repository used by the controllers
var Repo *Repository

// Instancia/Configura o repositório para todos os controllers
func NewRepository(a *config.AppConfig, db *driver.Database) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// Configura o repositório usado nos controllers
func SetRepository(r *Repository) {
	Repo = r
}

//#endregion

//#region: common endpoints
// default
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// remoteIP := r.RemoteAddr
	// log.Println("Home: Request incoming from: ", remoteIP)
}

//#endregion
