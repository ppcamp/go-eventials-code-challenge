// contains all postgres functions used to connect with database

package dbrepo

import (
	"context"
	"time"

	"yawoen.com/app/internal/dtos"
	"yawoen.com/app/internal/models"
)

//#region: Companies table
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
	m.App.InfoLog.Println("company -> FindById searching...")

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

func (m *postgresDBRepo) CompanyCreate(data *dtos.CompanyCreate) error {
	// if don't responde in 3 seconds, close it
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT
			INTO companies (name,addresszip, website)
			VALUES ($1,$2,$3);`

	m.App.InfoLog.Println("company -> Creating a new element...")

	_, err := m.DB.ExecContext(ctx, query, data.Name, data.AddressZip, data.Website)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) CompanyGet() ([]models.Company, error) {
	// if don't responde in 3 seconds, close it
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT *
		FROM
			companies;`

	m.App.InfoLog.Println("company -> GetAll searching...")

	var companies []models.Company
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return companies, err
	}

	// Itera adiciona cada linha da requisição no vetor
	for rows.Next() {
		var company models.Company
		err := rows.Scan(
			&company.Id,
			&company.Name,
			&company.AddressZip,
			&company.Website,
		)
		if err != nil {
			return companies, err
		}
		companies = append(companies, company)
	}
	m.App.InfoLog.Println("company -> GetAll sending response")

	return companies, nil
}

func (m *postgresDBRepo) CompanyPatch() ([]models.Company, error) { return nil, nil }

func (m *postgresDBRepo) CompanyPut() ([]models.Company, error) { return nil, nil }

func (m *postgresDBRepo) CompanyDelete() ([]models.Company, error) { return nil, nil }
