-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS withdrawal

(
    uuid            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user"          text,
    method          UUID REFERENCES payment_method (uuid),
    amount          DOUBLE PRECISION NOT NULL,
    details         TEXT,
    state           VARCHAR(20) CHECK (state IN ('PENDING', 'COMPLETED', 'REJECTED')),
    completion_date TIMESTAMP,
    rate            DOUBLE PRECISION,
    currency_source TEXT,
    currency_target TEXT,
    deleted         BOOLEAN          DEFAULT FALSE,
    created_at      TIMESTAMP(6),
    created_by      TEXT,
    updated_at      TIMESTAMP(6),
    updated_by      TEXT,
    deleted_at      TIMESTAMP(6),
    deleted_by      TEXT
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

