
DO $$ 
DECLARE 
 expected_rows integer := 49998;
 C integer := 0;

 actual_rows   record;
BEGIN 
    INSERT INTO instrument (
        symbol, 
        name, 
        description, 
        currency_id, 
        instrument_class_id, 
        from_date, thru_date, 
        created_at, created_by, 
        updated_at, updated_by
    ) SELECT 
        isin,
        bats_name,
        company_name,
        instrument.id,
        3,
        '2001-01-01', '2999-01-01', 
        current_timestamp, 'ETL', 
        current_timestamp, 'ETL' 
    FROM etl.etl_bats_instrument JOIN instrument ON (
        upper(etl_bats_instrument.currency) = instrument.symbol 
        AND instrument.instrument_class_id = 1)
    ;

    GET DIAGNOSTICS c = ROW_COUNT;

    RAISE INFO 'INSERT ROW COUNT %', c ;

    SELECT COUNT(id) INTO actual_rows FROM instrument;
   
    --ASSERT actual_rows.count = expected_rows,
    --    'Expect number of rows in instrument to be ' || expected_rows || ' not ' || actual_rows;
END $$;

