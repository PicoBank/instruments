-- +goose Up

CREATE TABLE institution (
    institution_id  	SERIAL,
	name				VARCHAR(80) NOT NULL,
	description			VARCHAR(255),
	acronym				VARCHAR(25),
	--
	from_date			TIMESTAMP(3) NOT NULL,
	thru_date			TIMESTAMP(3),
	--
	created_at			TIMESTAMP(3) NOT NULL DEFAULT NOW(),
	created_by			VARCHAR(25) NOT NULL,
	updated_at			TIMESTAMP(3) NOT NULL DEFAULT NOW(),
	updated_by			VARCHAR(25) NOT NULL,
	--
    CONSTRAINT pk_institution PRIMARY KEY (institution_id)
);

CREATE TABLE institution_role (
			institution_id		INTEGER NOT NULL,
			instrument_id		INTEGER NOT NULL,
			description			VARCHAR(255),
			--
			from_date			TIMESTAMP(3) NOT NULL,
			thru_date			TIMESTAMP(3),
			--
			created_at			TIMESTAMP(3) NOT NULL DEFAULT NOW(),
			created_by			VARCHAR(25) NOT NULL,
			updated_at			TIMESTAMP(3) NOT NULL DEFAULT NOW(),
			updated_by			VARCHAR(25) NOT NULL,
			--
			CONSTRAINT pk_institution_role PRIMARY KEY (institution_id, instrument_id),
			CONSTRAINT fk_institution_role_institution FOREIGN KEY (institution_id) REFERENCES institution (institution_id),
			CONSTRAINT fk_institution_role_instrument FOREIGN KEY (instrument_id) REFERENCES instrument (instrument_id)
);

-- +goose Down

DROP TABLE institution_role;
DROP TABLE institution;