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

// Obtem uma companhia com base no seu identificador
func (m *Repository) CompanyGetOne(w http.ResponseWriter, r *http.Request) {
	sid := chi.URLParam(r, "id")

	if !govalidator.IsInt(sid) {
		helpers.ClientError(w, 400, "The id should be an integer")
	}
	id, err := strconv.Atoi(sid)
	if err != nil {
		helpers.ServerError(w, err)
	}

	company, err := m.DB.CompanyFindById(id)
	if err != nil {
		helpers.ServerError(w, err)
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
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Cria uma nova companhia com base no seu identificador
func (m *Repository) CompaniesPost(w http.ResponseWriter, r *http.Request) {
	var company dtos.CompanyCreate
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// Tentando criar um elemento
	// err = m.DB.CompanyCreate(&company)
	// if err != nil {
	// 	helpers.ServerError(w, err)
	// }

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	// w.Write()
}

func (m *Repository) CompaniesGetMany(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("Is in Production? %t", m.App.InProduction)))
	companies, err := m.DB.CompanyGet()
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
	companies, err := m.DB.CompanyGet()
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
	companies, err := m.DB.CompanyGet()
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
	companies, err := m.DB.CompanyGet()
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
