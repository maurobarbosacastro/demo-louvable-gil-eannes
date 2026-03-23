-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

alter table category
add legacy_id text;

alter table country
add legacy_id text;

alter table partner
add legacy_id text;

alter table store
add legacy_id text;

alter table payment_method
add legacy_id text;

alter table withdrawal
add legacy_id text;

alter table reward
add legacy_id text;

alter table transaction
add legacy_id text;

alter table referral
add legacy_id text;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
