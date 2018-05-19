package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Instrument represents an instrument
type Instrument struct {
	ID          uint32 `orm:"pk;column(instrument_id)"`
	Symbol      string `faker:"word"`
	Name        string
	Description string           `json:",omitempty"`
	Class       *InstrumentClass `orm:"rel(fk);column(instrument_class_id)"`
	Currency    *Instrument      `orm:"null;rel(fk);column(currency_id)" json:",omitempty"`
	FromDate    time.Time
	ThruDate    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   string
	UpdatedBy   string
}

func init() {
	orm.RegisterModel(new(InstrumentClass))
}

// GetAllInstruments returns the list of all instruments
// TODO: this is obviously a very simplistic behavior...
func GetAllInstruments() (instruments []Instrument) {
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("instrument").RelatedSel()

	qs.All(&instruments)
	return
}

// GetOneInstrument returns one instrument from its ID
func GetOneInstrument(ID uint32) (instrument *Instrument, err error) {
	instrument = &Instrument{ID: ID}
	o := orm.NewOrm()
	o.Using("default")

	err = o.Read(instrument)
	if err != nil {
		return nil, err
	}
	return instrument, nil
}
