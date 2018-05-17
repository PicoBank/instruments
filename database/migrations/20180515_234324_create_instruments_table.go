package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateInstrumentsTable_20180515_234324 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateInstrumentsTable_20180515_234324{}
	m.Created = "20180515_234324"

	migration.Register("CreateInstrumentsTable_20180515_234324", m)
}

// Run the migrations
func (m *CreateInstrumentsTable_20180515_234324) Up() {
	m.SQL(`CREATE TABLE instruments.instruments (
		instrument_id  SERIAL,
		symbol         VARCHAR(25) NOT NULL,
		from_date      TIMESTAMP(3) NOT NULL,
		thru_date      TIMESTAMP(3),
		created_at     TIMESTAMP(3) NOT NULL DEFAULT NOW(),
		updated_at     TIMESTAMP(3) NOT NULL DEFAULT NOW(),
		created_by     VARCHAR(25) NOT NULL,
	
		CONSTRAINT pk_instrument PRIMARY KEY (instrument_id)
	)`)
}

// Reverse the migrations
func (m *CreateInstrumentsTable_20180515_234324) Down() {
	m.SQL(`DROP TABLE instruments.instruments`)
}
