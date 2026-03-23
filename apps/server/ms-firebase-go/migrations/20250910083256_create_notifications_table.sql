-- +goose Up
-- +goose StatementBegin
CREATE TYPE notification_state AS ENUM ('pending', 'delivered', 'canceled', 'error');

CREATE TABLE notification (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    date TIMESTAMP NOT NULL,
    state notification_state NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by TEXT,
    updated_at TIMESTAMP,
    updated_by TEXT
);

CREATE INDEX idx_notification_state ON notification(state);
CREATE INDEX idx_notification_date ON notification(date);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notification;
DROP TYPE IF EXISTS notification_state;
-- +goose StatementEnd
