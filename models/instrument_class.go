package models

// InstrumentClass represents an instrument class
type InstrumentClass struct {
	ID   uint32 `orm:"pk"`
	Name string
}
