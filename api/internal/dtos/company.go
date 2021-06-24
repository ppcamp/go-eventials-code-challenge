// contains the DTO that's used to check for these fields

package dtos

type CompanyCreate struct {
	Name       string
	AddressZip string
	Website    string

	Errors errors
}

// Valid returns true if there are no errors, otherwise false
func (f *CompanyCreate) Valid() bool {
	return len(f.Errors) == 0
}

// // New initializes a form struct
// func New(data CreateCompany) *CreateCompany {
// 	return &CreateCompany{
// 		data,
// 		errors(map[string][]string{}),
// 	}
// }

// // Required checks for required fields
// func (f *CreateCompany) Required(fields ...string) {
// 	for _, field := range fields {
// 		value := f.Get(field)
// 		if strings.TrimSpace(value) == "" {
// 			f.Errors.Add(field, "This field cannot be blank")
// 		}
// 	}
// }

// // Has checks if form field is in post and not empty
// func (f *CreateCompany) Has(field string) bool {
// 	x := f.Get(field)
// 	if x == "" {
// 		return false
// 	}
// 	return true
// }

// // MinLength checks for string minimum length
// func (f *CreateCompany) MinLength(field string, length int) bool {
// 	x := f.Get(field)
// 	if len(x) < length {
// 		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
// 		return false
// 	}
// 	return true
// }

// // IsEmail checks for valid email address
// func (f *CreateCompany) IsEmail(field string) {
// 	if !govalidator.IsEmail(f.Get(field)) {
// 		f.Errors.Add(field, "Invalid email address")
// 	}
// }
