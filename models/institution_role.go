package models

import (
	"time"
)

// Institution represents an institution
type InstitutionRole struct {
	InstitutionID uint32
	InstrumentID  uint32
	Description   string `json:",omitempty"`
	FromDate      time.Time
	ThruDate      time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
}

func init() {
	//orm.RegisterModel(new(InstitutionRole))
}
