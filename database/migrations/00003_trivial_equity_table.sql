-- +goose Up

CREATE TABLE equity
(
    instrument_id SERIAL,
    pay_currency_id INTEGER,
    exercise_currency_id INTEGER,
    company_currency_id INTEGER,
	--
	from_date			TIMESTAMP(3) NOT NULL,
	thru_date			TIMESTAMP(3),
	--
	created_at			TIMESTAMP(3) NOT NULL DEFAULT NOW(),
	created_by			VARCHAR(25) NOT NULL,
	updated_at			TIMESTAMP(3) NOT NULL DEFAULT NOW(),
	updated_by			VARCHAR(25) NOT NULL,
	--
    CONSTRAINT pk_equity PRIMARY KEY (instrument_id),
    CONSTRAINT fk_equity_instrument FOREIGN KEY (instrument_id) REFERENCES instrument (instrument_id),
    CONSTRAINT fk_equity_pay_currency FOREIGN KEY (pay_currency_id) REFERENCES instrument (instrument_id),
    CONSTRAINT fk_equity_exercise_currency FOREIGN KEY (exercise_currency_id) REFERENCES instrument (instrument_id),
    CONSTRAINT fk_equity_company_currency FOREIGN KEY (company_currency_id) REFERENCES instrument (instrument_id)
);

-- +goose Down

DROP TABLE equity;