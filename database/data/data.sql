\set ON_ERROR_STOP on

TRUNCATE instrument CASCADE;

-- Currency
INSERT INTO instrument (instrument_id, symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by, datum, audit) VALUES 
    (10001, 'USD', 'US Dollars'         , 'US Dollars'          , 1, null, '2010-01-01', '2099-12-31', 'init', 'init', '{"id":10001,"symbol":"USD","name":"US Dollars","description":"US Dollars"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (10002, 'EUR', 'Euro'               , 'Euro'                , 1, null, '2010-01-01', '2099-12-31', 'init', 'init', '{"id":10002,"symbol":"EUR","name":"Euro","description":"Euro"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (10003, 'CHF', 'Swiss Franc'        , 'Swiss Franc'         , 1, null, '2010-01-01', '2099-12-31', 'init', 'init', '{"id":10003,"symbol":"CHF","name":"Swiss Franc","description":"Swiss Franc"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (10004, 'AUD', 'Australian Dollar'  , 'Australian Dollar'   , 1, null, '2010-01-01', '2099-12-31', 'init', 'init', '{"id":10004,"symbol":"AUD","name":"Australian Dollars","description":"Australian Dollars"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (10005, 'CAD', 'Canadian Dollar'    , 'Canadian Dollar'     , 1, null, '2010-01-01', '2099-12-31', 'init', 'init', '{"id":10005,"symbol":"CAD","name":"Canadian Dollars","description":"Canadian Dollars"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (10006, 'GBP', 'UK Pound'           , 'UK Pound'            , 1, null, '2010-01-01', '2099-12-31', 'init', 'init', '{"id":10006,"symbol":"GBP","name":"OK Pound","description":"UK Pound"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}')
;

-- Bond
INSERT INTO instrument (instrument_id,symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by, datum, audit) VALUES 
    (20001, '0049804980', '6.657%LLOYDS BK PERP 144A'  , '6.657%LLOYDS BK PERP 144A'   , 2, 10001, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":20001,"symbol":"0049804980","name":"6.657%LLOYDS BK PERP 144A","description":"6.657%LLOYDS BK PERP 144A"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (20002, '0387368000', '4 5/8 JP MORGAN PERP CC'    , '4 5/8 JP MORGAN PERP CC'     , 2, 10006, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":20002,"symbol":"0387368000","name":"4 5/8 JP MORGAN PERP CC","description":"4 5/8 JP MORGAN PERP CC"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (20003, '0370143260', '6 3/4 CAIXABANK PERP'       , '6 3/4 CAIXABANK PERP'        , 2, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":20003,"symbol":"0370143260","name":"6 3/4 CAIXABANK PERP","description":"6 3/4 CAIXABANK PERP"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (20004, '0278496700', 'LA MONDIALE  5.11 %  VAR '  , 'LA MONDIALE  5.11 %  VAR '   , 2, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":20004,"symbol":"0278496700","name":"LA MONDIALE  5.11 %  VAR","description":"LA MONDIALE  5.11 %  VAR"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}')
;

-- Equity
INSERT INTO instrument (instrument_id,symbol, name, description, instrument_class_id, currency_id, from_date, thru_date, created_by, updated_by, datum, audit) VALUES 
    (30001, 'FR0010285965', 'ALMIL', '1000MERCIS'          , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":30001,"symbol":"FR0010285965","name":"ALMIL","description":"1000MERCIS"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30002, 'BE0974275076', 'VALOR', '2VALORISE'           , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":30002,"symbol":"BE0974275076","name":"VALOR","description":"2VALORISE"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30003, 'CH0008853209', 'AGTA' , 'AGTA RECORD'         , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":30003,"symbol":"CH0008853209","name":"AGTA","description":"AGTA RECORD"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30004, 'CA1125851040', 'BAMA' , 'BROOKFIELD ASSET M'  , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":30004,"symbol":"CA1125851040","name":"BAMA","description":"BROOKFIELD ASSET M"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30005, 'FR0000052896', 'CHAU' , 'CHAUF.URB.'          , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":30005,"symbol":"FR0000052896","name":"CHAU","description":"CHAUF.URB."}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30006, 'PTCDU0AE0003', 'CDU'  , 'CONDURIL'            , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":30006,"symbol":"PTCDU0AE0003","name":"CDU","description":"CONDURIL"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30007, 'NL0011509294', 'CURE' , 'CURETIS'             , 3, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"id":30007,"symbol":"NL0011509294","name":"CURE","description":"CURETIS"}', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}')
;

INSERT INTO Equity (instrument_id,pay_currency_id,exercise_currency_id,company_currency_id,from_date,thru_date,created_by,updated_by,audit) VALUES 
    (30001, 10002, 10002, 10002, '2018-01-23', '2099-12-31', 'init', 'init', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30002, 10002, 10002, 10001, '2018-01-23', '2099-12-31', 'init', 'init', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30003, 10002, 10002, 10003, '2018-01-23', '2099-12-31', 'init', 'init', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30004, 10002, 10002, 10004, '2018-01-23', '2099-12-31', 'init', 'init', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30005, 10002, 10002, 10005, '2018-01-23', '2099-12-31', 'init', 'init', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30006, 10002, 10002, 10001, '2018-01-23', '2099-12-31', 'init', 'init', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}'),
    (30007, 10002, 10002, 10006, '2018-01-23', '2099-12-31', 'init', 'init', '{"createdAt":"2010-01-01", "createdBy":"init","updatedAt":"2010-01-01", "updatedBy":"init"}')
;