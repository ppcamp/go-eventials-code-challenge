package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"yawoen.com/app/internal/dtos"
)

func TestCompanyPost(t *testing.T) {
	// make our handler a http.HandlerFunc
	handler := http.HandlerFunc(Repo.CompanyPost)

	log.Println("Test")
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
		jsonPayload, _ := json.MarshalIndent(data, "", " ")                // convert to json payload
		request := bytes.NewBuffer(jsonPayload)                            // convert the json to bytes handler
		req, _ := http.NewRequest("POST", "/search-availability", request) // create our request
		req.Header.Set("Content-Type", "application/json")                 // set the request header
		rr := httptest.NewRecorder()                                       // create our response recorder-> requirement for http.ResponseWriter
		handler.ServeHTTP(rr, req)                                         // make the request to our handler
		if rr.Code != test.ExpectedStatus {
			t.Errorf("[POST] company/ Got status %d expected %d", rr.Code, test.ExpectedStatus)
		}
	}

}
