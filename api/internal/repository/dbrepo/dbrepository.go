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

//#region: Criando o padrão para repositórios postgres
// Cria uma nova conexão com o postgres
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepository {

	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

//#endregion
