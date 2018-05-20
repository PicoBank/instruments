package models

import (
	"github.com/astaxie/beego/orm"
)

// InstrumentClass represents an instrument class
type InstrumentClass struct {
	ID   uint32 `orm:"pk;column(instrument_class_id)"`
	Name string
}

func init() {
	orm.RegisterModel(new(InstrumentClass))
}

// GetAllInstrumentClasses lists all instrument classes
func GetAllInstrumentClasses() (instrumentClasses []InstrumentClass) {
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("instrument_class")

	qs.All(&instrumentClasses)
	return
}
