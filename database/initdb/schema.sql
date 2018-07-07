\set ON_ERROR_STOP on

REVOKE ALL ON DATABASE picobank FROM public;
CREATE USER instruments PASSWORD 'raspberry';
GRANT CONNECT ON DATABASE picobank TO instruments;

\connect picobank pi

CREATE SCHEMA instruments AUTHORIZATION instruments;

-- The following is better to ensure all tools run in the proper schema
ALTER ROLE instruments IN DATABASE picobank SET search_path = instruments;
