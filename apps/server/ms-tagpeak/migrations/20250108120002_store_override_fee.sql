-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE store ADD COLUMN override_fee FLOAT NULL;

ALTER TABLE partner DROP COLUMN commission_rate;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
