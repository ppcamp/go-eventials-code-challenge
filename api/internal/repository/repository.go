package repository

import (
	"yawoen.com/app/internal/dtos"
	"yawoen.com/app/internal/models"
)

// Essa interface vai conter todas as rotas usadas para conectar no banco
type DatabaseRepository interface {
	CompanyCreate(*dtos.CompanyCreate) error                     // create element
	CompanyFindById(int) (models.Company, error)                 // search for this id
	CompanyFindByQuery(...interface{}) ([]models.Company, error) // update all fields
	CompanyEditAll(*dtos.CompanyCreate) error                    // fetch data
	CompanyEditByQuery(...interface{}) error                     // update specific fields
	CompanyDelete(int) error                                     // delete element by id
}
