package models

import (
	"time"
)

// InstitutionRole represents an institution role
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
