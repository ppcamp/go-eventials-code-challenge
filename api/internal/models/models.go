// Contém as entidades (tabelas do banco)
package models

// Companies table
type Company struct {
	Id         int
	Name       string
	AddressZip NullString
	Website    NullString
}
