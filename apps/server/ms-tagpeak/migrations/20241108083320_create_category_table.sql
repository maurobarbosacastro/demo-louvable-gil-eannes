-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS category (
    uuid       UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       TEXT,
    created_at TIMESTAMP(6),
    created_by TEXT,
    updated_at TIMESTAMP(6),
    updated_by TEXT,
    deleted    BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMP(6),
    deleted_by TEXT
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
