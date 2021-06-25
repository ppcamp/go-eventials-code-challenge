package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"yawoen.com/app/internal/dtos"
)

func TestCompanyPost(t *testing.T) {
	// make our handler a http.HandlerFunc
	handler := http.HandlerFunc(Repo.CompanyPost)

	// configure tests
	var tests = []struct {
		TestName       string
		ExpectedStatus int
		Name           string
		AddressZip     string
		Website        string
	}{
		{"successfully-created", 201, "SomeUser", "", ""},
		{"validation-problem", 400, "", "", ""},
	}

	for _, test := range tests {
		data := dtos.CompanyCreate{
			Name:       test.Name,
			Website:    test.Website,
			AddressZip: test.AddressZip,
		}

		/// change it
		jsonPayload, _ := json.MarshalIndent(data, "", " ")    // convert to json payload
		request := bytes.NewBuffer(jsonPayload)                // convert the json to bytes handler
		req, _ := http.NewRequest("POST", "/company", request) // create our request
		req.Header.Set("Content-Type", "application/json")     // set the request header
		rr := httptest.NewRecorder()                           // create our response recorder-> requirement for http.ResponseWriter
		handler.ServeHTTP(rr, req)                             // make the request to our handler
		if rr.Code != test.ExpectedStatus {
			t.Errorf("[POST] /company Got status %d expected %d", rr.Code, test.ExpectedStatus)
		}
	}

}

func TestCompanyGet(t *testing.T) {
	// make our handler a http.HandlerFunc
	handler := http.HandlerFunc(Repo.CompanyGetMany)

	// configure tests
	var tests = []struct {
		TestName       string
		ExpectedStatus int
		Name           string
		AddressZip     string
		Website        string
	}{
		{"searching-by-name", 200, "SomeUser", "", ""},
		{"get-all-elements", 200, "", "", ""},
		{"validation-field-problem", 400, "", "123456", "1235125"},
	}

	for _, test := range tests {
		data := map[string]string{
			"name":    test.Name,
			"website": test.Website,
			"zip":     test.AddressZip,
		}

		/// change it
		req, _ := http.NewRequest("GET", "/company", nil) // create our request
		q := req.URL.Query()
		for k, v := range data {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		req.Header.Set("Content-Type", "application/json") // set the request header
		rr := httptest.NewRecorder()                       // create our response recorder-> requirement for http.ResponseWriter
		handler.ServeHTTP(rr, req)                         // make the request to our handler
		if rr.Code != test.ExpectedStatus {
			t.Errorf("[GET] /company Got status %d expected %d", rr.Code, test.ExpectedStatus)
		}
	}

}

func TestCompanyPut(t *testing.T) {
	routes := setUpRoutes()
	mockServer := httptest.NewTLSServer(routes)

	defer mockServer.Close()

	// configure tests
	var tests = []struct {
		TestName       string
		ExpectedStatus int
		Name           string
		AddressZip     string
		Website        string
	}{
		{"successfully-edited", 200, "SomeUser", "", ""},
		{"validation-problem", 400, "", "", ""},
	}

	for _, test := range tests {
		data := dtos.CompanyCreate{
			Name:       test.Name,
			Website:    test.Website,
			AddressZip: test.AddressZip,
		}

		// change it
		jsonPayload, _ := json.MarshalIndent(data, "", " ")                    // convert to json payload
		request := bytes.NewBuffer(jsonPayload)                                // convert the json to bytes handler
		req, _ := http.NewRequest("PUT", mockServer.URL+"/company/1", request) // create our request
		req.Header.Set("Content-Type", "application/json")                     // set the request header
		resp, _ := mockServer.Client().Do(req)                                 // makes the request handling with id paths
		if resp.StatusCode != test.ExpectedStatus {
			t.Errorf("[PUT] company/ Got status %d expected %d", resp.StatusCode, test.ExpectedStatus)
		}
	}

}

func TestCompanyGetOne(t *testing.T) {
	routes := setUpRoutes()
	mockServer := httptest.NewTLSServer(routes)

	defer mockServer.Close()

	// configure tests
	var tests = []struct {
		TestName       string
		ExpectedStatus int
		Id             string
	}{
		{"successfully-edited", 200, "1"},
		{"validation-problem", 400, "asd"},
	}

	for _, test := range tests {
		// change it
		req, _ := http.NewRequest("GET", mockServer.URL+"/company/"+test.Id, nil) // create our request
		req.Header.Set("Content-Type", "application/json")                        // set the request header
		resp, _ := mockServer.Client().Do(req)                                    // makes the request handling with id paths
		if resp.StatusCode != test.ExpectedStatus {
			t.Errorf("[GET] company/ Got status %d expected %d", resp.StatusCode, test.ExpectedStatus)
		}
	}
}

func TestCompanyDelete(t *testing.T) {
	routes := setUpRoutes()
	mockServer := httptest.NewTLSServer(routes)

	defer mockServer.Close()

	// configure tests
	var tests = []struct {
		TestName       string
		ExpectedStatus int
		Id             string
	}{
		{"successfully-edited", 200, "1"},
		{"validation-problem", 400, "asd"},
	}

	for _, test := range tests {
		// change it
		req, _ := http.NewRequest("DELETE", mockServer.URL+"/company/"+test.Id, nil) // create our request
		req.Header.Set("Content-Type", "application/json")                           // set the request header
		resp, _ := mockServer.Client().Do(req)                                       // makes the request handling with id paths
		if resp.StatusCode != test.ExpectedStatus {
			t.Errorf("[DELETE] company/ Got status %d expected %d", resp.StatusCode, test.ExpectedStatus)
		}
	}
}
