-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS withdrawal_audit
(
    uuid            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    withdrawal_id   UUID REFERENCES withdrawal (uuid),
    action          VARCHAR(20) CHECK (action IN ('ADD', 'UPDATE', 'DELETE')),
    "user"          TEXT,
    method          TEXT,
    amount          DOUBLE PRECISION,
    details         TEXT,
    state           VARCHAR(20) CHECK (state IN ('PENDING', 'COMPLETED', 'REJECTED')),
    completion_date TIMESTAMP,
    rate            DOUBLE PRECISION,
    currency_source TEXT,
    currency_target TEXT,
    deleted         BOOLEAN,
    created_at      TIMESTAMP(6),
    created_by      text,
    updated_at      TIMESTAMP(6),
    updated_by      text,
    deleted_at      TIMESTAMP(6),
    deleted_by      text
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
