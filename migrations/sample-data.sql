\set ON_ERROR_STOP on

\connect picobank pi;

INSERT INTO instruments.instruments (symbol, from_date, created_by) VALUES 
    ('NYC', NOW(), 'test'),
    ('NZU', NOW(), 'test'),
    ('KPO', NOW(), 'test');