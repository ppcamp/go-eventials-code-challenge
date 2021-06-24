package main

import "testing"

func TestSetUp(t *testing.T) {
	_, err := setUp()
	if err != nil {
		t.Error("Falhou no test de configurações")
	}
}
