-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE user_camunda_process RENAME TO camunda_process;
ALTER TABLE camunda_process RENAME COLUMN user_uuid to field_uuid;
ALTER TABLE camunda_process DROP COLUMN job_type;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
