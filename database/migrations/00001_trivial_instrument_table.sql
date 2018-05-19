-- +goose Up
CREATE TABLE instrument (
		instrument_id  SERIAL,
		symbol         VARCHAR(25) NOT NULL,
		name           VARCHAR(80) NOT NULL,
		description    VARCHAR(255),
		from_date      TIMESTAMP(3) NOT NULL,
		thru_date      TIMESTAMP(3),
		created_at     TIMESTAMP(3) NOT NULL DEFAULT NOW(),
		updated_at     TIMESTAMP(3) NOT NULL DEFAULT NOW(),
		created_by     VARCHAR(25) NOT NULL,
	
		CONSTRAINT pk_instrument PRIMARY KEY (instrument_id)
	);
	
-- +goose Down
DROP TABLE instrument;