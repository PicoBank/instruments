\set ON_ERROR_STOP on

TRUNCATE instrument CASCADE;

-- Currency
INSERT INTO instrument (instrument_id, symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by) VALUES 
    (10001, 'USD', 'US Dollars'         , 'US Dollars'          , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (10002, 'EUR', 'Euro'               , 'Euro'                , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (10003, 'CHF', 'Swiss Franc'        , 'Swiss Franc'         , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (10004, 'AUD', 'Australian Dollar'  , 'Australian Dollar'   , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (10005, 'CAD', 'Canadian Dollar'    , 'Canadian Dollar'     , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (10006, 'GBP', 'UK Pound'           , 'UK Pound'            , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' )
;

-- Bond
INSERT INTO instrument (instrument_id,symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by) VALUES 
    (20001, '0049804980', '6.657%LLOYDS BK PERP 144A'  , '6.657%LLOYDS BK PERP 144A'   , 2, 10001, '2018-01-23', '2099-12-31', 'init', 'init'),
    (20002, '0387368000', '4 5/8 JP MORGAN PERP CC'    , '4 5/8 JP MORGAN PERP CC'     , 2, 10006, '2018-01-23', '2099-12-31', 'init', 'init'),
    (20003, '0370143260', '6 3/4 CAIXABANK PERP'       , '6 3/4 CAIXABANK PERP'        , 2, 10002, '2018-01-23', '2099-12-31', 'init', 'init'),
    (20004, '0278496700', 'LA MONDIALE  5.11 %  VAR '  , 'LA MONDIALE  5.11 %  VAR '   , 2, 10002, '2018-01-23', '2099-12-31', 'init', 'init')
;

-- Equity
INSERT INTO instrument (instrument_id,symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by) VALUES 
    (30001, 'FR0010285965', 'ALMIL', '1000MERCIS'          , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30002, 'BE0974275076', 'VALOR', '2VALORISE'           , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30003, 'CH0008853209', 'AGTA' , 'AGTA RECORD'         , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30004, 'CA1125851040', 'BAMA' , 'BROOKFIELD ASSET M'  , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30005, 'FR0000052896', 'CHAU' , 'CHAUF.URB.'          , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30006, 'PTCDU0AE0003', 'CDU'  , 'CONDURIL'            , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30007, 'NL0011509294', 'CURE' , 'CURETIS'             , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init')
;

INSERT INTO Equity (instrument_id,pay_currency_id,exercise_currency_id,company_currency_id,from_date,thru_date,created_by,updated_by) VALUES 
    (30001, 10002, 10002, 10002, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30002, 10002, 10002, 10001, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30003, 10002, 10002, 10003, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30004, 10002, 10002, 10004, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30005, 10002, 10002, 10005, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30006, 10002, 10002, 10001, '2018-01-23', '2099-12-31', 'init', 'init'),
    (30007, 10002, 10002, 10006, '2018-01-23', '2099-12-31', 'init', 'init')
;