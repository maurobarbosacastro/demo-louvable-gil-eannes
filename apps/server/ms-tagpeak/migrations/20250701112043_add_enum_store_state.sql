-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd


ALTER TABLE store
ADD CONSTRAINT check_store_state
    CHECK (state IN ('ACTIVE', 'INACTIVE', 'PENDING', 'BLOCKED'));

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
