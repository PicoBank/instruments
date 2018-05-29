-- +goose Up
CREATE TABLE instrument_class (
	instrument_class_id DECIMAL(2) NOT NULL,
	name                VARCHAR(16),
	--
	CONSTRAINT pk_instrument_class PRIMARY KEY (instrument_class_id)
);

INSERT INTO instrument_class (instrument_class_id, name) VALUES 
    (1,   'Currency'),
    (2,   'Bond'),
    (3,   'Equity'),
    (4,   'Fund'),
    (5,   'Future'),
    (6,   'Option'),
    (7,   'Entitlement'),
    (8,   'Index'),
    (9,   'InterestRate'),
    (10,  'Commodity'),
    (11,  'Miscellaneous');


CREATE TABLE instrument (
    instrument_id               SERIAL,
    symbol         				VARCHAR(25) NOT NULL,
    name           				VARCHAR(80) NOT NULL,
    description    				VARCHAR(255),
    instrument_class_id         DECIMAL(2) NOT NULL,
    currency_id                 INTEGER,
    from_date                   TIMESTAMP(3) NOT NULL,
    thru_date                   TIMESTAMP(3),
    created_at                  TIMESTAMP(3) NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMP(3) NOT NULL DEFAULT NOW(),
    created_by                  VARCHAR(25) NOT NULL,
    updated_by                  VARCHAR(25) NOT NULL,

    CONSTRAINT pk_instrument PRIMARY KEY (instrument_id),
    CONSTRAINT fk_instrument_class FOREIGN KEY (instrument_class_id) REFERENCES instrument_class (instrument_class_id),
    CONSTRAINT fk_instrument_currency FOREIGN KEY (currency_id) REFERENCES instrument (instrument_id)
);

-- +goose Down
DROP TABLE instrument;
DROP TABLE instrument_class;