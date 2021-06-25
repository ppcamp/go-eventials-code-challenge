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
		helpers.ClientError(w, http.StatusNotFound, "Not found any element with this id")
		return
	}

	// Manually parsing solving marshal problems with not null
	var aux = struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		AddressZip string `json:"zip"`
		Website    string `json:"website"`
	}{company.Id, company.Name, "", ""}
	if company.AddressZip.Valid {
		aux.AddressZip = company.AddressZip.String
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

// Find elements basing on query parameters
//
// Return 200 for success 500 If something unexpected occurred
func (m *Repository) CompanyGetMany(w http.ResponseWriter, r *http.Request) {
	// creating the query struct
	company := dtos.CompanyQuery{
		Name:       r.URL.Query().Get("name"),
		AddressZip: r.URL.Query().Get("zip"),
		Website:    r.URL.Query().Get("website"),
	}

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
			return
		}

		// Return the errors messages
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	companies, err := m.DB.CompanyFindByQuery(&company)
	if err != nil {
		helpers.ServerError(w, err)
	}

	// converte para json
	out, err := json.MarshalIndent(companies, "", " ")
	if err != nil {
		helpers.ServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Creates a new company element
//
// It takes a dtos.CompanyCreate element as parameter. And makes validations on it
// Return 400 for validation errors and 201 for success
func (m *Repository) CompanyPost(w http.ResponseWriter, r *http.Request) {
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
			return
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

	// Otherwise, just sends the created status flag (default, not necessary)
	w.WriteHeader(http.StatusCreated)
}

// Delete an element basing on its id
//
// Return 200 for success 500 If something unexpected occurred
func (m *Repository) CompanyDelete(w http.ResponseWriter, r *http.Request) {
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

	err = m.DB.CompanyDelete(id)
	if err != nil {
		helpers.ClientError(w, http.StatusNotFound, "Not found any element with this id")
		return
	}

	// Otherwise, sent ok (default)
}

// Edit all fields for a given id
//
// Return 200 for success 500 If something unexpected occurred
func (m *Repository) CompanyPut(w http.ResponseWriter, r *http.Request) {
	// get id
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

	var company dtos.CompanyCreate
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err = json.NewDecoder(r.Body).Decode(&company)
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
			return
		}

		// Return the errors messages
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	err = m.DB.CompanyEditAll(id, &company)
	if err != nil {
		helpers.ClientError(w, 400, err.Error())
		return
	}

	// Otherwise, just sends the created status flag (default, not necessary)
}

// TODO: implement this controller
func (m *Repository) CompanyPatch(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("Not implemented yet"))
}

// Edit all fields for a given id
//
// Return 200 for success 500 If something unexpected occurred
func (m *Repository) CompanyPutWebsite(w http.ResponseWriter, r *http.Request) {
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
			return
		}

		// Return the errors messages
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	modifiedId, err := m.DB.CompanyEditWebsite(&company)
	if err != nil {
		helpers.ClientError(w, 400, err.Error())
		return
	}

	var response = struct {
		Modified int `json:"modifiedId"`
	}{Modified: modifiedId}

	// converte para json
	out, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// Otherwise, just sends the ok status flag (default, not necessary)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
