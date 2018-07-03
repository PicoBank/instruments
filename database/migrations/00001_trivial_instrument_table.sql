-- +goose Up
CREATE TABLE instrument_class (
	id      SMALLINT NOT NULL,
	name    VARCHAR(16),
	--
	CONSTRAINT pk_instrument_class PRIMARY KEY (id)
);

INSERT INTO instrument_class (id, name) VALUES 
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

CREATE TABLE instrument_source (
	id      SMALLINT NOT NULL,
	name    VARCHAR(16),
	--
	CONSTRAINT pk_instrument_source PRIMARY KEY (id)
);

INSERT INTO instrument_source (id, name) VALUES 
    (0,   'MISC'),
    (1,   'BATS/EU'),
    (2,   'CSI/CASH');

CREATE TABLE instrument (
    id                          SERIAL,
    symbol         				VARCHAR(25) NOT NULL,
    name           				VARCHAR(80) NOT NULL,
    description    				VARCHAR(255),
    instrument_class_id         SMALLINT NOT NULL,
    instrument_source_id        SMALLINT NOT NULL,
    currency_id                 INTEGER,
    from_date                   TIMESTAMP(3) NOT NULL,
    thru_date                   TIMESTAMP(3),
    created_at                  TIMESTAMP(3) NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMP(3) NOT NULL DEFAULT NOW(),
    created_by                  VARCHAR(25) NOT NULL,
    updated_by                  VARCHAR(25) NOT NULL,

    CONSTRAINT pk_instrument PRIMARY KEY (id),
    CONSTRAINT fk_instrument_class FOREIGN KEY (instrument_class_id) REFERENCES instrument_class (id),
    CONSTRAINT fk_instrument_source FOREIGN KEY (instrument_source_id) REFERENCES instrument_source (id),
    CONSTRAINT fk_instrument_currency FOREIGN KEY (currency_id) REFERENCES instrument (id)
);

CREATE INDEX idx_instrument_instrument_class_id on instrument(instrument_class_id);
CREATE INDEX idx_instrument_instrument_source_id on instrument(instrument_source_id);
CREATE INDEX idx_instrument_instrument_currency on instrument(currency_id);

-- on laisse 10000 ids pour les instruments 'de référence': currencies
ALTER SEQUENCE instrument_id_seq RESTART WITH 10000;

-- +goose Down
DROP TABLE instrument;
DROP TABLE instrument_source;
DROP TABLE instrument_class;