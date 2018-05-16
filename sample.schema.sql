\set ON_ERROR_STOP on

\connect postgres
DROP DATABASE IF EXISTS picobank;
DROP OWNED BY pi;
DROP USER IF EXISTS pi;

CREATE USER pi PASSWORD 'raspberry';
ALTER ROLE pi SUPERUSER;
ALTER ROLE pi SET search_path TO instruments;
CREATE DATABASE picobank OWNER pi;
\connect picobank pi;
CREATE SCHEMA instruments;

CREATE TABLE instruments.instruments (
    code character varying,
    value numeric
);

GRANT SELECT ON TABLE instruments.instruments TO pi;

INSERT INTO instruments.instruments VALUES ('NYC', 1.34);
INSERT INTO instruments.instruments VALUES ('NZU', 3.14);

