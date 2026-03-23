-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_member_transaction_cash_reward', 'Member Cash rewards from transactions', '0', true, 'text', now(),
        'system');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
