package dbrepo

import (
	"database/sql"

	"yawoen.com/app/internal/config"
	"yawoen.com/app/internal/repository"
)

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// Creates a test repository
func NewTestRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepository {
	return &testDBRepo{
		App: a,
		DB:  conn,
	}
}
