-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE country
    ADD CONSTRAINT abbreviation_unique UNIQUE (abbreviation);

CREATE TABLE IF NOT EXISTS store_country
(
    country_code TEXT NOT NULL,
    store_uuid   UUID NOT NULL,

    PRIMARY KEY (country_code, store_uuid),
    CONSTRAINT fk_country FOREIGN KEY (country_code) REFERENCES country (abbreviation) ON DELETE CASCADE,
    CONSTRAINT fk_store FOREIGN KEY (store_uuid) REFERENCES store (uuid) ON DELETE CASCADE
);

ALTER TABLE store
    DROP COLUMN country_uuid;


SELECT 'down SQL query';
-- +goose StatementEnd
