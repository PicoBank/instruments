package models

import (
	"time"
)

// Instrument represents an instrument
type Instrument struct {
	ID           uint32 `orm:"pk;column(id)"`
	Symbol       string `faker:"word"`
	Name         string
	Description  string           `json:",omitempty"`
	Class        *InstrumentClass `orm:"rel(fk);column(instrument_class_id)"`
	Currency     *Instrument      `orm:"null;rel(fk);column(currency_id)" json:",omitempty"`
	Institutions []*Institution   `orm:"rel(m2m);rel_table(institution_role)"`
	FromDate     time.Time
	ThruDate     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreatedBy    string
	UpdatedBy    string
}
