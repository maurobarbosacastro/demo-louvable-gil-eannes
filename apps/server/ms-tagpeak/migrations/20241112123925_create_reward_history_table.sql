-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS reward_history
(
    uuid        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    reward_uuid UUID,
    rate        DOUBLE PRECISION,
    units       DOUBLE PRECISION,
    cash_reward DOUBLE PRECISION,
    deleted     BOOLEAN          DEFAULT FALSE,
    created_at  TIMESTAMP(6),
    created_by  TEXT,
    updated_at  TIMESTAMP(6),
    updated_by  TEXT,
    deleted_at  TIMESTAMP(6),
    deleted_by  TEXT,
    FOREIGN KEY (reward_uuid) REFERENCES reward (uuid)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
