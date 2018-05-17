\set ON_ERROR_STOP on

CREATE USER pi PASSWORD 'raspberry';
ALTER ROLE pi SUPERUSER;
--ALTER ROLE pi SET search_path TO instruments;
CREATE DATABASE picobank OWNER pi;

\connect picobank pi;

CREATE SCHEMA instruments;

CREATE TABLE instruments.instruments (
    code character varying,
    value numeric
);

GRANT SELECT ON TABLE instruments.instruments TO pi;