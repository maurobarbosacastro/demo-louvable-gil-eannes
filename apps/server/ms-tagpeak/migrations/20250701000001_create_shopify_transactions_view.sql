-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE OR REPLACE VIEW shopify_actionable_transactions AS
SELECT t.uuid as transaction_uuid,
    s.uuid as store_uuid,
    s.name as store_name,
    t.state as transaction_state,
    t.order_date,
    s.average_reward_activation_time,
    CASE
        WHEN s.average_reward_activation_time IS NULL
            THEN 0
        ELSE
            CASE
                -- If activation time is already passed, return 0
                WHEN (CURRENT_DATE - t.order_date) > CAST(split_part(s.average_reward_activation_time, ' ', 1) AS INTEGER)
                    THEN 0
                -- Otherwise, calculate remaining time (+1 to fix off-by-one)
                ELSE CAST(split_part(s.average_reward_activation_time, ' ', 1) AS INTEGER)
                    - (CURRENT_DATE - t.order_date)
                    + 1
                END
        END as remaining_time
FROM transaction t
         JOIN store s ON t.store_uuid = s.uuid
         JOIN partner p ON s.affiliate_partner_code = p.code
WHERE p.code = 'shopify' AND upper(t.state) = 'TRACKED' AND t.deleted = false;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
