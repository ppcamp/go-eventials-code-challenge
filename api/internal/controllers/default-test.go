package controllers

import (
	"yawoen.com/app/internal/config"
	"yawoen.com/app/internal/repository/dbrepo"
)

// NewTestRepo creates a new repository
func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewTestingsRepo(a),
	}
}
