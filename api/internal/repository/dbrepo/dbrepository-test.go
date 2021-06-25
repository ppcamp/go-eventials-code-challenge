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

// creates a new testing repository
func NewTestingsRepo(a *config.AppConfig) repository.DatabaseRepository {
	return &testDBRepo{
		App: a,
	}
}
