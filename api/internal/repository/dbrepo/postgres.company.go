// contains all postgres functions used to connect with database

package dbrepo

import (
	"context"
	"time"

	"yawoen.com/app/internal/dtos"
	"yawoen.com/app/internal/models"
)

// Find an element in the database based on its id
//
// Returns the element that matched or error
func (m *postgresDBRepo) CompanyFindById(id int) (models.Company, error) {
	// if don't responde in 3 seconds, close it
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT *
		FROM
			companies
		WHERE
			id = $1;`

	var company models.Company
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&company.Id,
		&company.Name,
		&company.AddressZip,
		&company.Website,
	)
	if err != nil {
		return company, err
	}
	return company, nil
}

// Create an new element based on the dto.CompanyCreate element
//
// Returns nil for success
func (m *postgresDBRepo) CompanyCreate(data *dtos.CompanyCreate) error {
	// if don't responde in 3 seconds, close it
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT
			INTO companies (name,addresszip, website)
			VALUES ($1,$2,$3);`

	_, err := m.DB.ExecContext(ctx, query, data.Name, data.AddressZip, data.Website)
	if err != nil {
		return err
	}
	return nil
}

// Searching using query parameters
//
// Returns all elements that matches with this
func (m *postgresDBRepo) CompanyEditByQuery(...interface{}) error {
	return nil
}

//
func (m *postgresDBRepo) CompanyEditAll(data *dtos.CompanyCreate) error {
	return nil
}

//
func (m *postgresDBRepo) CompanyFindByQuery(...interface{}) ([]models.Company, error) {
	return nil, nil
}

func (m *postgresDBRepo) CompanyDelete(id int) error {
	return nil
}
