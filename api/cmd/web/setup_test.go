package main

import (
	"os"
	"testing"

	"github.com/subosito/gotenv"
)

func TestMain(tests *testing.M) {
	// Configures the environment values, used to fetch
	gotenv.Load("./../../.env")
	os.Exit(tests.Run())
}

// type myHandler struct{}

// func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// }
