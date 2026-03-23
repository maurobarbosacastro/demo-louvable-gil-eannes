-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS user_camunda_process
(
    uuid                    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_uuid               UUID,
    process_instance_key    BIGINT,
    process_id              TEXT,
    job_type                TEXT
    );

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
