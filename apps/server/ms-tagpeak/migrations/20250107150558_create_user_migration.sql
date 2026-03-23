-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS user_migration (
    id SERIAL PRIMARY KEY,
    user_uuid UUID NOT NULL,
    legacy_id INTEGER NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
