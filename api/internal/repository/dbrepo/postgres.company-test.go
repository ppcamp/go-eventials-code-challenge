// contains all postgres functions used to connect with database
// TODO: Refactor the code to add more error treatment

package dbrepo

import (
	"context"
	"time"

	"yawoen.com/app/internal/dtos"
	"yawoen.com/app/internal/models"
)

// Find an element in the database based on its id
//
// Returns the element that matched or error on failure
func (m *testDBRepo) CompanyFindById(id int) (models.Company, error) {
	// if don't responde in 3 seconds, close it
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, name, addresszip, website
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
// Returns nil for success or error on failure
func (m *testDBRepo) CompanyCreate(data *dtos.CompanyCreate) error {
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

// Find an element basing on the query passed through
//
// If the field is not empty, add them to query search
// Return nil if occurred some system error
func (m *testDBRepo) CompanyFindByQuery(data *dtos.CompanyQuery) ([]models.Company, error) {
	// if don't responde in 3 seconds, close it
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// NOTE: the query string don't treat the escape symbol
	query := `
		SELECT id, name, addresszip, website
			FROM companies
			WHERE ($1='' OR name ILIKE $2)
				AND ($3='' OR addresszip ILIKE $4)
				AND ($5='' OR website ILIKE $6);`

	var companies []models.Company
	rows, err := m.DB.QueryContext(
		ctx,
		query,
		data.Name,
		"%"+data.Name+"%",
		data.AddressZip,
		"%"+data.AddressZip+"%",
		data.Website,
		"%"+data.Website+"%",
	)
	if err != nil {
		return companies, err
	}

	for rows.Next() {
		var row models.Company
		err := rows.Scan(
			&row.Id,
			&row.Name,
			&row.AddressZip,
			&row.Website)
		if err != nil {
			return companies, err
		}
		companies = append(companies, row)
	}

	return companies, nil
}

// Delete an element based on its id
//
// Return nil on success or error on failure
func (m *testDBRepo) CompanyDelete(id int) error {
	// if don't responde in 3 seconds, close it
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Check if exists
	_, err := m.CompanyFindById(id)
	if err != nil {
		return err
	}

	// if exists, remove it
	query := `DELETE FROM companies WHERE id=$1;`
	_, err = m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// Edit all fields for this element that matches with this id
//
// Returns nil on success or error on failure
func (m *testDBRepo) CompanyEditAll(id int, data *dtos.CompanyCreate) error {
	// if don't responde in 3 seconds, close it
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Check if exists
	_, err := m.CompanyFindById(id)
	if err != nil {
		return err
	}

	query := `
		UPDATE companies
			SET name=$2,addresszip=$3, website=$4
			WHERE id=$1;`

	// Run the query
	_, err = m.DB.ExecContext(ctx, query, id, data.Name, data.AddressZip, data.Website)
	if err != nil {
		return err
	}

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
	// if don't responde in 3 seconds, close it
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		UPDATE companies
			SET website=$1
			WHERE
				(name ILIKE $2) AND (addresszip ILIKE $3)
			RETURNING id;`

	// Run the query
	var id int
	err := m.DB.QueryRowContext(
		ctx,
		query,
		data.Website,
		"%"+data.Name+"%",
		"%"+data.AddressZip+"%").Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}
