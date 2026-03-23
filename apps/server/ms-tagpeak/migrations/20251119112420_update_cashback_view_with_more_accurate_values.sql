-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

drop view if exists cashback;

CREATE OR REPLACE VIEW cashback AS
SELECT t.uuid               AS transaction_uuid,
    t."user"             AS "user",
    sv.uuid              AS store_visit_uuid,
    s.uuid               AS store_uuid,
    s.name               AS store_name,
    s.logo               AS store_logo,
    s.percentage_cashout AS store_percentage_cashout,
    s.cashback_value     AS store_cashback_value,
    s.cashback_type      AS store_cashback_type,
    t.order_date         AS date,
    t.amount_source,
    t.amount_target,
    t.amount_user,
    t.currency_source,
    t.currency_target,
    CASE
        WHEN t.manual_commission IS NOT NULL THEN t.manual_commission
        ELSE t.commission_target
        END              AS network_commission,
    CASE
        WHEN t.state::text = 'VALIDATED'::text AND r.uuid IS NOT NULL THEN r.state
        ELSE t.state
        END              AS status,
    t.cashback,
    r.uuid               AS reward_uuid,
    r.isin,
    r.current_reward_source,
    r.current_reward_target,
    r.current_reward_user,
    r.initial_price,
    r.title,
    r.end_date,
    r.created_at         AS start_date,
    r.origin,
    r.stopped_at,
    sv.reference         AS store_visit_reference
FROM transaction t
         LEFT JOIN reward r ON t.uuid = r.transaction_uuid AND r."user" = t."user"
         LEFT JOIN store_visit sv ON t.store_visit_uuid = sv.uuid
         LEFT JOIN store s ON t.store_uuid = s.uuid
WHERE t.deleted = false
    -- Include this row if: there's a reward for this user OR there ARE rewards for other users (so we need to show the transaction without reward)
    AND (r.uuid IS NOT NULL OR EXISTS (
    SELECT 1 FROM reward r2
    WHERE r2.transaction_uuid = t.uuid
        AND r2."user" != t."user"
))

UNION

SELECT t.uuid               AS transaction_uuid,
    r."user"             AS "user",
    sv.uuid              AS store_visit_uuid,
    s.uuid               AS store_uuid,
    s.name               AS store_name,
    s.logo               AS store_logo,
    s.percentage_cashout AS store_percentage_cashout,
    s.cashback_value     AS store_cashback_value,
    s.cashback_type      AS store_cashback_type,
    t.order_date         AS date,
    t.amount_source,
    t.amount_target,
    t.amount_user,
    t.currency_source,
    t.currency_target,
    CASE
        WHEN t.manual_commission IS NOT NULL THEN t.manual_commission
        ELSE t.commission_target
        END              AS network_commission,
    r.state              AS status,
    t.cashback,
    r.uuid               AS reward_uuid,
    r.isin,
    r.current_reward_source,
    r.current_reward_target,
    r.current_reward_user,
    r.initial_price,
    r.title,
    r.end_date,
    r.created_at         AS start_date,
    r.origin,
    r.stopped_at,
    sv.reference         AS store_visit_reference
FROM transaction t
         INNER JOIN reward r ON t.uuid = r.transaction_uuid AND r."user" != t."user"
         LEFT JOIN store_visit sv ON t.store_visit_uuid = sv.uuid
         LEFT JOIN store s ON t.store_uuid = s.uuid
WHERE t.deleted = false

UNION

-- Add transactions with no rewards at all
SELECT t.uuid               AS transaction_uuid,
    t."user"             AS "user",
    sv.uuid              AS store_visit_uuid,
    s.uuid               AS store_uuid,
    s.name               AS store_name,
    s.logo               AS store_logo,
    s.percentage_cashout AS store_percentage_cashout,
    s.cashback_value     AS store_cashback_value,
    s.cashback_type      AS store_cashback_type,
    t.order_date         AS date,
    t.amount_source,
    t.amount_target,
    t.amount_user,
    t.currency_source,
    t.currency_target,
    CASE
        WHEN t.manual_commission IS NOT NULL THEN t.manual_commission
        ELSE t.commission_target
        END              AS network_commission,
    t.state              AS status,
    t.cashback,
    NULL                 AS reward_uuid,
    NULL                 AS isin,
    NULL                 AS current_reward_source,
    NULL                 AS current_reward_target,
    NULL                 AS current_reward_user,
    NULL                 AS initial_price,
    NULL                 AS title,
    NULL                 AS end_date,
    NULL                 AS start_date,
    NULL                 AS origin,
    NULL                 AS stopped_at,
    NULL        AS store_visit_reference
FROM transaction t
         LEFT JOIN store_visit sv ON t.store_visit_uuid = sv.uuid
         LEFT JOIN store s ON t.store_uuid = s.uuid
WHERE t.deleted = false
    AND NOT EXISTS (
    SELECT 1 FROM reward r
    WHERE r.transaction_uuid = t.uuid
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP VIEW IF EXISTS cashback;

CREATE OR REPLACE VIEW cashback
        (transaction_uuid, "user", store_visit_uuid, store_uuid, store_name, store_logo, store_percentage_cashout,
         store_cashback_value, store_cashback_type, date, amount_source, amount_target, amount_user,
         currency_source, currency_target, network_commission, status, cashback, reward_uuid, isin,
         current_reward_source, current_reward_target, current_reward_user, initial_price, title, end_date,
         start_date, origin, stopped_at)
as
SELECT t.uuid               AS transaction_uuid,
    CASE
        WHEN r.uuid IS NOT NULL THEN r."user"
        ELSE t."user"
        END              AS "user",
    sv.uuid              AS store_visit_uuid,
    s.uuid               AS store_uuid,
    s.name               AS store_name,
    s.logo               AS store_logo,
    s.percentage_cashout AS store_percentage_cashout,
    s.cashback_value     AS store_cashback_value,
    s.cashback_type      AS store_cashback_type,
    t.order_date         AS date,
    t.amount_source,
    t.amount_target,
    t.amount_user,
    t.currency_source,
    t.currency_target,
    CASE
        WHEN t.manual_commission IS NOT NULL THEN t.manual_commission
        ELSE t.commission_target
        END              AS network_commission,
    CASE
        WHEN t.state::text = 'VALIDATED'::text AND r.uuid IS NOT NULL THEN r.state
        ELSE t.state
        END              AS status,
    t.cashback,
    r.uuid               AS reward_uuid,
    r.isin,
    r.current_reward_source,
    r.current_reward_target,
    r.current_reward_user,
    r.initial_price,
    r.title,
    r.end_date,
    r.created_at         AS start_date,
    r.origin,
    r.stopped_at
FROM transaction t
         LEFT JOIN reward r ON t.uuid = r.transaction_uuid
         LEFT JOIN store_visit sv ON t.store_visit_uuid = sv.uuid
         LEFT JOIN store s ON t.store_uuid = s.uuid
WHERE t.deleted = false;
