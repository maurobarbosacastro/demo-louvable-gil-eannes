-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS payment_method

(
    uuid       UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       TEXT NOT NULL,
    code       TEXT NOT NULL,
    created_at TIMESTAMP(6),
    created_by VARCHAR(255),
    updated_at TIMESTAMP(6),
    updated_by VARCHAR(255),
    deleted    BOOLEAN          DEFAULT FALSE,
    deleted_at TIMESTAMP(6),
    deleted_by TEXT
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
