package models

import (
	"github.com/astaxie/beego/orm"
)

type Instrument struct {
	Code  string `orm:"pk"`
	Value float64
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
