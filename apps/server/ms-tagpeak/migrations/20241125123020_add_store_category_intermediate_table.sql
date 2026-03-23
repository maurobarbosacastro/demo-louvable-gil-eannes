-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';


CREATE TABLE IF NOT EXISTS store_category
(
    category_code TEXT NOT NULL,
    store_uuid    UUID NOT NULL,

    PRIMARY KEY (category_code, store_uuid),
    CONSTRAINT fk_language FOREIGN KEY (category_code) REFERENCES category (code) ON DELETE CASCADE,
    CONSTRAINT fk_store FOREIGN KEY (store_uuid) REFERENCES store (uuid) ON DELETE CASCADE
);

ALTER TABLE store
    DROP COLUMN category_uuid;


SELECT 'down SQL query';
-- +goose StatementEnd
