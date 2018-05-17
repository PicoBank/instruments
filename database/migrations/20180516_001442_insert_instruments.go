package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertInstruments_20180516_001442 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertInstruments_20180516_001442{}
	m.Created = "20180516_001442"

	migration.Register("InsertInstruments_20180516_001442", m)
}

// Run the migrations
func (m *InsertInstruments_20180516_001442) Up() {
	m.SQL(`INSERT INTO instruments.instruments (symbol, from_date, created_by) VALUES 
		('NYC', NOW(), 'test'),
    	('NZU', NOW(), 'test'),
    	('KPO', NOW(), 'test')`)
}

// Reverse the migrations
func (m *InsertInstruments_20180516_001442) Down() {
	m.SQL(`DELETE FROM instruments.instruments`)
}
