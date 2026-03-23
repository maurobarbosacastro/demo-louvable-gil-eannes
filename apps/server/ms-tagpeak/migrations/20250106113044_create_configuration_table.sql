-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS configuration
(
    id
    serial
    primary
    key,
    code
    text
    not
    null,
    name
    text
    not
    null,
    value
    text
    not
    null,
    editable
    boolean
    default
    true
    not
    null,
    data_type
    text
    not
    null,
    created_at
    TIMESTAMP
(
    6
),
    created_by TEXT,
    updated_at TIMESTAMP
(
    6
),
    updated_by TEXT,
    deleted BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMP
(
    6
),
    deleted_by TEXT
    );


-- Features
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('expired', 'Expired flag feature', 'false', true, 'bool', now(), 'system');

-- Tagpeak configurations
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('commission_rate', 'Default Commission Rate', '4', true, 'number', now(), 'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('transaction_fixed_fee', 'Tagpeak Investment fixed fee', '2', true, 'number', now(), 'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('tagpeak_default_currency', 'Tagpeak Default Currency', 'EUR', true, 'string', now(), 'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('withdrawal_balance_minimum', 'Minimum Withdrawal amount', '20', true, 'number', now(), 'system');

-- Referrals configurations
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_default_status', 'Default status for Referrer on Signup', 'Silver', true, 'text', now(), 'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_member_friend_cash_reward', 'Member Friend Cash Reward', '0', true, 'number', now(), 'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_silver_status_goal', 'Number of referral for Silver status', '1', true, 'number', now(), 'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_silver_status_goal_amount', 'Amount to get Silver status', '500', true, 'text', now(), 'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_silver_transaction_cash_reward', 'Silver Cash rewards from transactions', '10', true, 'text', now(),
        'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_silver_friend_reward_share', 'Silver Share from friends cash reward', '5', true, 'text', now(),
        'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_gold_status_goal', 'Number of referral for Gold status', '5', true, 'number', now(), 'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_gold_status_goal_amount', 'Amount to get Gold status', '5000', true, 'text', now(), 'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_gold_transaction_cash_reward', 'Gold Cash rewards from transactions', '20', true, 'text', now(),
        'system');
insert into configuration (code, name, value, editable, data_type, created_at, created_by)
values ('referral_gold_friend_reward_share', 'Gold Share from friends cash reward', '10', true, 'text', now(),
        'system');


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
