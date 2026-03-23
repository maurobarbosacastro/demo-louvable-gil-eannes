-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE reward
    ADD COLUMN title TEXT;

ALTER TABLE reward
    ADD COLUMN details TEXT;

ALTER TABLE reward_audit
    ADD COLUMN title TEXT;

ALTER TABLE reward_audit
    ADD COLUMN details TEXT;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
