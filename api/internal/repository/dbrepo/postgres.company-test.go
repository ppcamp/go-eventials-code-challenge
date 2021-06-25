// contains all postgres functions used to connect with database
// TODO: Refactor the code to add more error treatment

package dbrepo

import (
	"yawoen.com/app/internal/dtos"
	"yawoen.com/app/internal/models"
)

// Find an element in the database based on its id
//
// Returns the element that matched or error on failure
func (m *testDBRepo) CompanyFindById(id int) (models.Company, error) {
	var company models.Company

	return company, nil
}

// Create an new element based on the dto.CompanyCreate element
//
// Returns nil for success or error on failure
func (m *testDBRepo) CompanyCreate(data *dtos.CompanyCreate) error {
	return nil
}

// Find an element basing on the query passed through
//
// If the field is not empty, add them to query search
// Return nil if occurred some system error
func (m *testDBRepo) CompanyFindByQuery(data *dtos.CompanyQuery) ([]models.Company, error) {
	var companies []models.Company
	return companies, nil
}

// Delete an element based on its id
//
// Return nil on success or error on failure
func (m *testDBRepo) CompanyDelete(id int) error {
	return nil
}

// Edit all fields for this element that matches with this id
//
// Returns nil on success or error on failure
func (m *testDBRepo) CompanyEditAll(id int, data *dtos.CompanyCreate) error {
	return nil
}

// Updating only the requested fields
//
// Return nil on success
// TODO: implement this function
func (m *testDBRepo) CompanyEditByQuery(id int, data ...interface{}) error {
	return nil
}

// Edit website for this element that matches with the `name` and `zip`
//
// Returns the modified row
// Returns (id,nil) on success or (-1,error) on failure
func (m *testDBRepo) CompanyEditWebsite(data *dtos.CompanyCreate) (int, error) {
	return 3, nil
}
