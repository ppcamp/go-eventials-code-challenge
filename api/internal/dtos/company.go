// contains the DTO that's used to check for these fields

package dtos

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

//#region: CreateDTO

type CompanyCreate struct {
	Name       string
	AddressZip string
	Website    string

	// Variable used to get all encountered errors
	Errors errors
}

// Start the map used to track all possible errors
func (f *CompanyCreate) ErrorsInit() {
	f.Errors = map[string][]string{}
}

// Check for errors
// Valid returns true if there are no errors, otherwise false
func (f *CompanyCreate) Valid() bool {
	// checking for required fields
	if len(f.Name) == 0 {
		f.Errors.Add("Name", "Is a required field and should not be empty")
	}

	// checking for max lenghts
	var fields = []struct {
		Name    string
		Size    int
		Maxsize int
	}{
		{"Name", len(f.Name), 200},
		{"AddressZip", len(f.AddressZip), 5},
		{"Website", len(f.Website), 200},
	}
	for _, i := range fields {
		if i.Size > i.Maxsize {
			f.Errors.Add(i.Name, fmt.Sprintf("Accept at maximum %d chars", i.Maxsize))
		}
	}

	if !govalidator.IsInt(f.AddressZip) && len(f.AddressZip) > 0 {
		f.Errors.Add("AddressZip", "Should be only digits")
	}

	if !govalidator.IsURL(f.Website) && len(f.Website) > 0 {
		f.Errors.Add("Website", "Should be an url formatted string")
	}

	return len(f.Errors) == 0
}

//#endregion

//#region: QueryDTO
type CompanyQuery struct {
	Name       string
	AddressZip string
	Website    string

	// Variable used to get all encountered errors
	Errors errors
}

// Start the map used to track all possible errors
func (f *CompanyQuery) ErrorsInit() {
	f.Errors = map[string][]string{}
}

func (f *CompanyQuery) Valid() bool {
	// checking for max lenghts
	var fields = []struct {
		Name    string
		Size    int
		Maxsize int
	}{
		{"Name", len(f.Name), 200},
		{"AddressZip", len(f.AddressZip), 5},
		{"Website", len(f.Website), 200},
	}
	for _, i := range fields {
		if i.Size > i.Maxsize {
			f.Errors.Add(i.Name, fmt.Sprintf("Accept at maximum %d chars", i.Maxsize))
		}
	}

	if !govalidator.IsInt(f.AddressZip) && len(f.AddressZip) > 0 {
		f.Errors.Add("AddressZip", "Should be only digits")
	}

	if !govalidator.IsURL(f.Website) && len(f.Website) > 0 {
		f.Errors.Add("Website", "Should be an url formatted string")
	}

	return len(f.Errors) == 0
}

//#endregion
