-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_token (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_uuid TEXT NOT NULL,
    token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_at TIMESTAMP,
    updated_by TEXT,
    UNIQUE(user_uuid, token)
);

CREATE INDEX idx_user_token_user_uuid ON user_token(user_uuid);
CREATE INDEX idx_user_token_token ON user_token(token);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_token;
-- +goose StatementEnd
