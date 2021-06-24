package repository

import (
	"yawoen.com/app/internal/dtos"
	"yawoen.com/app/internal/models"
)

// Essa interface vai conter todas as rotas usadas para conectar no banco
type DatabaseRepository interface {
	CompanyFindById(int) (models.Company, error) // search for this id
	CompanyCreate(*dtos.CompanyCreate) error     // create element
	CompanyGet() ([]models.Company, error)       // fetch data
	CompanyPatch() ([]models.Company, error)     // update specific fields
	CompanyPut() ([]models.Company, error)       // update all fields
	CompanyDelete() ([]models.Company, error)    // delete element by id
}
