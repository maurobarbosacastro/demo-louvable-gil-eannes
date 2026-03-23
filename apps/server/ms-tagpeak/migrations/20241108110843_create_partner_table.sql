-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd


CREATE TABLE IF NOT EXISTS partner
(
    uuid                 UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name                 TEXT NOT NULL,
    e_commerce_platform  TEXT,
    commission_rate      FLOAT,
    validation_period    INT,
    deep_link            TEXT,
    deep_link_identifier TEXT,
    sub_identifier       TEXT,
    percentage_tagpeak   FLOAT,
    percentage_invested  FLOAT,
    created_at           TIMESTAMP(6),
    created_by           TEXT,
    updated_at           TIMESTAMP(6),
    updated_by          TEXT,
    deleted              BOOLEAN DEFAULT FALSE,
    deleted_at           TIMESTAMP(6),
    deleted_by           TEXT
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
