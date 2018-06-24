\set ON_ERROR_STOP on

TRUNCATE instrument CASCADE;

-- Currency
INSERT INTO instrument (id, symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by, datum, audit) VALUES 
    (1, 'USD', 'US Dollars'         , 'US Dollars'          , 1, null, '2010-01-01', '2099-12-31', 'init', 'init', '{"id":10001,"symbol":"USD","name":"US Dollars","description":"US Dollars","identifiers": [{"scheme":"default","value":"USD"},{"scheme":"AA","value":"USD"},{"scheme":"BB","value":"USD"},{"scheme":"CC","value":"USD"},{"scheme":"DD","value":"USD"},{"scheme":"DD","value":"USD"}]}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (2, 'EUR', 'Euro'               , 'Euro'                , 1, null, '2010-01-01', '2099-12-31', 'init', 'init', '{"id":10002,"symbol":"EUR","name":"Euro","description":"Euro","identifiers": [{"scheme":"default","value":"EUR"},{"scheme":"AA","value":"EUR"},{"scheme":"BB","value":"EUR"},{"scheme":"CC","value":"EUR"},{"scheme":"DD","value":"EUR"},{"scheme":"DD","value":"EUR"}]}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (3, 'CHF', 'Swiss Franc'        , 'Swiss Franc'         , 1, null, '2010-01-01', '2099-12-31', 'init', 'init', '{"id":10003,"symbol":"CHF","name":"Swiss Franc","description":"Swiss Franc","identifiers": [{"scheme":"default","value":"CHF"},{"scheme":"AA","value":"CHF"},{"scheme":"BB","value":"CHF"},{"scheme":"CC","value":"CHF"},{"scheme":"DD","value":"CHF"},{"scheme":"DD","value":"CHF"}]}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}')
;