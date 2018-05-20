package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Institution represents an institution
type Institution struct {
	ID          uint32 `orm:"pk;column(institution_id)"`
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

func init() {
	orm.RegisterModel(new(Institution))
}

// GetAllInstitutions returns the list of all institutions
// TODO: this is obviously a very simplistic behavior...
func GetAllInstitutions() (institutions []Institution) {
	o := orm.NewOrm()
	o.Using("default")

	qs := o.QueryTable("institution").RelatedSel()

	qs.All(&institutions)
	return
}

// GetOneInstitution returns one institution from its ID
func GetOneInstitution(ID uint32) (institution *Institution, err error) {
	institution = &Institution{ID: ID}
	o := orm.NewOrm()
	o.Using("default")

	err = o.Read(institution)
	if err != nil {
		return nil, err
	}
	return institution, nil
}
