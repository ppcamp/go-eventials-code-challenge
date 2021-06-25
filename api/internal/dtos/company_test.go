package dtos

import (
	"fmt"
	"testing"
)

func TestCompanyCreateDTO(t *testing.T) {
	var tests = []struct {
		TestName      string
		ExpectedValid bool
		Name          string
		AddressZip    string
		Website       string
	}{
		{"missing-required-field", false, "", "", ""},
		{"max-length-size", false, "SomeOrgName", "123456", ""},
		{"website-isnt-url", false, "SomeOrgName", "12346", "someWrongData"},
		{"is-valid", true, "SomeOrgName", "12346", "www.google.com"},
		{"is-valid-empty-zip", true, "SomeOrgName", "", "www.google.com"},
		{"is-valid-empty-website", true, "SomeOrgName", "12345", ""},
		{"is-valid-empty-zip-and-website", true, "SomeOrgName", "", ""}}

	for _, test := range tests {
		data := CompanyCreate{
			Name:       test.Name,
			Website:    test.Website,
			AddressZip: test.AddressZip,
		}
		data.ErrorsInit()
		result := data.Valid()
		if result != test.ExpectedValid {
			t.Error(fmt.Sprintf("%s expect %t got %t", test.TestName, result, test.ExpectedValid))
		}
	}
}

func TestCompanyQueryDTO(t *testing.T) {
	var tests = []struct {
		TestName      string
		ExpectedValid bool
		Name          string
		AddressZip    string
		Website       string
	}{
		{"max-length-size", false, "SomeOrgName", "123456", ""},
		{"website-isnt-url", false, "SomeOrgName", "12346", "someWrongData"},
		{"is-valid", true, "SomeOrgName", "12346", "www.google.com"},
		{"is-valid-empty-zip", true, "SomeOrgName", "", "www.google.com"},
		{"is-valid-empty-website", true, "SomeOrgName", "12345", ""},
		{"is-valid-empty-zip-and-website", true, "SomeOrgName", "", ""}}

	for _, test := range tests {
		data := CompanyCreate{
			Name:       test.Name,
			Website:    test.Website,
			AddressZip: test.AddressZip,
		}
		data.ErrorsInit()
		result := data.Valid()
		if result != test.ExpectedValid {
			t.Error(fmt.Sprintf("%s expect %t got %t", test.TestName, result, test.ExpectedValid))
		}
	}
}
