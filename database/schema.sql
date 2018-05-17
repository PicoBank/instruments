\set ON_ERROR_STOP on

REVOKE ALL ON DATABASE picobank FROM public;
CREATE USER instruments PASSWORD 'raspberry';
GRANT CONNECT ON DATABASE picobank TO instruments;

\connect picobank

CREATE SCHEMA instruments AUTHORIZATION instruments;
SET search_path = instruments;
ALTER ROLE instruments IN DATABASE picobank SET search_path = instruments;
