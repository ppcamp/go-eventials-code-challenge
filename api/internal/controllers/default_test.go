package controllers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"yawoen.com/app/internal/driver"
)

//#region: tests

// methods that don't expect specific fields/validations
var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{{"home", "/", "GET", http.StatusOK}}

//#endregion

// Create a sample test repository
func TestNewRepo(t *testing.T) {
	var db driver.Database
	testRepo := NewRepository(&app, &db)

	if reflect.TypeOf(testRepo).String() != "*controllers.Repository" {
		t.Errorf("Did not get correct type from NewRepo: got %s, wanted *Repository", reflect.TypeOf(testRepo).String())
	}
}

// test endpoints that are global (don't depend of payload)
func TestHandlers(t *testing.T) {
	routes := setUpRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}
}
