-- +goose Up
-- +goose StatementBegin
CREATE TABLE topic (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_at TIMESTAMP,
    updated_by TEXT
);

CREATE INDEX idx_topic_name ON topic(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS topic;
-- +goose StatementEnd
