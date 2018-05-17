\set ON_ERROR_STOP on

\connect postgres
DROP DATABASE IF EXISTS picobank;
DROP OWNED BY pi;
DROP USER IF EXISTS pi;