-- +goose Up
CREATE TABLE instrument_json (
    id                          SERIAL,
    name           			    VARCHAR(80) NOT NULL,
    description    				VARCHAR(255),
    datum                       JSONB,
    audit                       JSONB,
    currency_id                 INTEGER,
    from_date                   TIMESTAMP(3) NOT NULL,
    thru_date                   TIMESTAMP(3),
    created_at                  TIMESTAMP(3) NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMP(3) NOT NULL DEFAULT NOW(),
    created_by                  VARCHAR(25) NOT NULL,
    updated_by                  VARCHAR(25) NOT NULL,

    CONSTRAINT pk_instrument_json PRIMARY KEY (id),
    CONSTRAINT fk_instrument_currency_json FOREIGN KEY (currency_id) REFERENCES instrument (id)
);

create index instrumentgin on instrument_json using gin (datum);

-- +goose Down
DROP TABLE instrument_json;

