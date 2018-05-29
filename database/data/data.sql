
// Currency
INSERT INTO instrument (instrument_id, symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by) VALUES 
    (70001, 'USD', 'US Dollars'         , 'US Dollars'          , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (70002, 'EUR', 'Euro'               , 'Euro'                , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (70003, 'CHF', 'Swiss Franc'        , 'Swiss Franc'         , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (70004, 'AUD', 'Australian Dollar'  , 'Australian Dollar'   , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (70005, 'CAD', 'Canadian Dollar'    , 'Canadian Dollar'     , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' ),
    (70006, 'GBP', 'UK Pound'           , 'UK Pound'            , 1, null, '2010-01-01', '2099-12-31', 'init', 'init' )
;

// Bond
INSERT INTO instrument (symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by) VALUES 
    ('0049804980', '6.657%LLOYDS BK PERP 144A'  , '6.657%LLOYDS BK PERP 144A'   , 2, 70001, '2018-01-23', '2099-12-31', 'init', 'init'),
    ('0387368000', '4 5/8 JP MORGAN PERP CC'    , '4 5/8 JP MORGAN PERP CC'     , 2, 70006, '2018-01-23', '2099-12-31', 'init', 'init'),
    ('0370143260', '6 3/4 CAIXABANK PERP'       , '6 3/4 CAIXABANK PERP'        , 2, 70002, '2018-01-23', '2099-12-31', 'init', 'init'),
    ('0278496700', 'LA MONDIALE  5.11 %  VAR '  , 'LA MONDIALE  5.11 %  VAR '   , 2, 70002, '2018-01-23', '2099-12-31', 'init', 'init')
;

// Equity
INSERT INTO instrument (symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by) VALUES 
    ('FR0010285965', 'ALMIL', '1000MERCIS'          , 3, 70002, '2018-01-23', '2099-12-31', 'init', 'init'),
    ('BE0974275076', 'VALOR', '2VALORISE'           , 3, 70002, '2018-01-23', '2099-12-31', 'init', 'init'),
    ('CH0008853209', 'AGTA' , 'AGTA RECORD'         , 3, 70002, '2018-01-23', '2099-12-31', 'init', 'init'),
    ('CA1125851040', 'BAMA' , 'BROOKFIELD ASSET M'  , 3, 70002, '2018-01-23', '2099-12-31', 'init', 'init'),
    ('FR0000052896', 'CHAU' , 'CHAUF.URB.'          , 3, 70002, '2018-01-23', '2099-12-31', 'init', 'init'),
    ('PTCDU0AE0003', 'CDU'  , 'CONDURIL'            , 3, 70002, '2018-01-23', '2099-12-31', 'init', 'init'),
    ('NL0011509294', 'CURE' , 'CURETIS'             , 3, 70002, '2018-01-23', '2099-12-31', 'init', 'init')
;
