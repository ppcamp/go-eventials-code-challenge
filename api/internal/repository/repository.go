package repository

import (
	"yawoen.com/app/internal/dtos"
	"yawoen.com/app/internal/models"
)

// it contains all routes used to connect with database
type DatabaseRepository interface {

	//#region: Company services
	CompanyCreate(*dtos.CompanyCreate) error                         // create element
	CompanyFindById(int) (models.Company, error)                     // search for this id
	CompanyFindByQuery(*dtos.CompanyQuery) ([]models.Company, error) // update all fields
	CompanyEditAll(int, *dtos.CompanyCreate) error                   // fetch data
	CompanyDelete(int) error                                         // delete element by id
	CompanyEditByQuery(int, ...interface{}) error                    // TODO: create a route to update specific fields
	CompanyEditWebsite(*dtos.CompanyCreate) (int, error)             //
	//#endregion

}
