package models

import (
	"database/sql"
	"encoding/json"
)

//#region: NullString type

type NullString struct {
	sql.NullString
}

// Scan implements the Scanner interface for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

//#endregion
