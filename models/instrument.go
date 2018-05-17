package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Instrument struct {
	InstrumentId string `orm:"pk"`
	Symbol       string
	FromDate     time.Time
	ThruDate     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	created_by   string
}

func init() {
	orm.RegisterModel(new(Instrument))
}

func GetAllInstruments() (instruments []Instrument) {
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("instruments").Limit(10)

	qs.All(&instruments)
	return
}
func (u *Instrument) TableName() string {
	return "instruments"
}
