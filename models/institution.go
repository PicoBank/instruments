package models

import (
	"time"
)

// Institution represents an institution
type Institution struct {
	ID          uint32 `orm:"pk;column(id)"`
	Name        string
	Description string        `json:",omitempty"`
	Acronym     string        `json:",omitempty"`
	Instruments []*Instrument `orm:"reverse(many)"`
	FromDate    time.Time
	ThruDate    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   string
	UpdatedBy   string
}
