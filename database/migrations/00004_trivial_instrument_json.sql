-- +goose Up

ALTER TABLE instrument
add column datum jsonb;

ALTER TABLE instrument
add column audit jsonb;

ALTER TABLE equity
add column audit jsonb;

-- +goose Down
ALTER TABLE instrument
DROP COLUMN datum;

ALTER TABLE instrument
DROP COLUMN audit;

ALTER TABLE equity
DROP COLUMN audit;
