package dbrepo

import (
	"database/sql"

	"yawoen.com/app/internal/config"
	"yawoen.com/app/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// Creating a Repository that will be used to hold all postgres connections
//
// Returns a postgres repository
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepository {

	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
