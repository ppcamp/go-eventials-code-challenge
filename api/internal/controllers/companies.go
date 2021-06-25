package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"
	"yawoen.com/app/internal/dtos"
	"yawoen.com/app/internal/helpers"
)

// Get a company based on its id
//
// Receives an id from url. The id should be an integer
// Return 400 for validation errors and 200 for success
func (m *Repository) CompanyGetOne(w http.ResponseWriter, r *http.Request) {
	sid := chi.URLParam(r, "id")

	if !govalidator.IsInt(sid) {
		helpers.ClientError(w, 400, "The id should be an integer")
		return
	}
	id, err := strconv.Atoi(sid)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	company, err := m.DB.CompanyFindById(id)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// Manually parsing solving marshal problems with not null
	var aux = struct {
		Id        int
		Name      string
		AdressZip string
		Website   string
	}{company.Id, company.Name, "", ""}
	if company.AddressZip.Valid {
		aux.AdressZip = company.AddressZip.String
	}
	if company.Website.Valid {
		aux.Website = company.Website.String
	}

	// converte para json
	out, err := json.MarshalIndent(aux, "", " ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Creates a new company element
//
// It takes a dtos.CompanyCreate element as parameter. And makes validations on it
// Return 400 for validation errors and 201 for success
// TODO: what happens when I don't pass a value?
func (m *Repository) CompaniesPost(w http.ResponseWriter, r *http.Request) {
	var company dtos.CompanyCreate

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// Check for errors
	company.ErrorsInit()
	if !company.Valid() {
		var error = struct {
			StatusCode int
			Message    dtos.Errors
		}{http.StatusBadRequest, dtos.Errors(company.Errors)}

		// convert to json
		out, err := json.MarshalIndent(error, "", " ")
		if err != nil {
			helpers.ServerError(w, err)
		}

		// Return the errors messages
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	err = m.DB.CompanyCreate(&company)
	if err != nil {
		helpers.ClientError(w, 400, err.Error())
		return
	}

	// Otherwise, just sends the created status flag
	// w.WriteHeader(http.StatusCreated)
	// w.Header().Set("Content-Type", "application/json")
}

func (m *Repository) CompaniesGetMany(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("Is in Production? %t", m.App.InProduction)))
	companies, err := m.DB.CompanyFindByQuery()
	if err != nil {
		helpers.ServerError(w, err)
	}

	// converte para json
	out, err := json.MarshalIndent(companies, "", " ")
	if err != nil {
		helpers.ServerError(w, err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) CompaniesPut(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("Is in Production? %t", m.App.InProduction)))
	companies, err := m.DB.CompanyFindByQuery()
	if err != nil {
		helpers.ServerError(w, err)
	}

	// converte para json
	out, err := json.MarshalIndent(companies, "", " ")
	if err != nil {
		helpers.ServerError(w, err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) CompaniesPatch(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("Is in Production? %t", m.App.InProduction)))
	companies, err := m.DB.CompanyFindByQuery()
	if err != nil {
		helpers.ServerError(w, err)
	}

	// converte para json
	out, err := json.MarshalIndent(companies, "", " ")
	if err != nil {
		helpers.ServerError(w, err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) CompaniesDelete(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("Is in Production? %t", m.App.InProduction)))
	companies, err := m.DB.CompanyFindByQuery()
	if err != nil {
		helpers.ServerError(w, err)
	}

	// converte para json
	out, err := json.MarshalIndent(companies, "", " ")
	if err != nil {
		helpers.ServerError(w, err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
