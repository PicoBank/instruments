package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Instrument struct {
	Id          uint32 `orm:"pk;column(instrument_id)"`
	Symbol      string
	Name        string
	Description string
	FromDate    time.Time
	ThruDate    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	created_by  string
}

func init() {
	orm.RegisterModel(new(Instrument))
}

func GetAllInstruments() (instruments []Instrument) {
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("instrument")

	qs.All(&instruments)
	return
}

func GetOneInstrument(Id uint32) (instrument *Instrument, err error) {
	instrument = &Instrument{Id: Id}
	o := orm.NewOrm()
	o.Using("default")

	err = o.Read(instrument)
	if err != nil {
		return nil, err
	}
	return instrument, nil
}
