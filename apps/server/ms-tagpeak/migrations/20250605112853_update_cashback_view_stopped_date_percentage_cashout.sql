-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

drop view public_stores;

CREATE OR REPLACE VIEW public_stores AS
SELECT s.uuid                                      as store_uuid,
    s.name,
    s.logo,
    s.state,
    s.created_at,
    COALESCE(visit_stats.total_visits, 0)       as total_visits,
    COALESCE(cashback_stats.total_cashbacks, 0) as total_cashbacks,
    array_agg(DISTINCT country.abbreviation)    as countries_code,
    array_agg(DISTINCT c.code)                  as categories_code
FROM store as s
         LEFT JOIN store_country ON store_country.store_uuid = s.uuid
         LEFT JOIN country on country.abbreviation = store_country.country_code
         LEFT JOIN store_category ON store_category.store_uuid = s.uuid
         LEFT JOIN public.category c on c.code = store_category.category_code
         LEFT JOIN (SELECT store_uuid, COUNT(*) as total_visits
                    FROM store_visit
                    GROUP BY store_uuid) visit_stats ON visit_stats.store_uuid = s.uuid
         LEFT JOIN (SELECT store_uuid, COUNT(*) as total_cashbacks
                    FROM transaction
                    GROUP BY store_uuid) cashback_stats ON cashback_stats.store_uuid = s.uuid
WHERE LOWER(s.state) = LOWER('ACTIVE')
    AND s.deleted = false
GROUP BY s.uuid, s.name, s.logo, s.state, visit_stats.total_visits, cashback_stats.total_cashbacks;

drop view cashback;

CREATE or replace VIEW cashback AS
SELECT
    t.uuid as transaction_uuid,
    case
        when r.uuid is not null then r."user"
        else t."user"
        end as "user",
    sv.uuid as store_visit_uuid,
    s.uuid as store_uuid,
    s.name as store_name,
    s.logo as store_logo,
    s.percentage_cashout as store_percentage_cashout,
    s.cashback_value as store_cashback_value,
    s.cashback_type as store_cashback_type,

    t.order_date as date,
    t.amount_source,
    t.amount_target,
    t.amount_user,
    t.currency_source,
    t.currency_target,
    CASE
        when t.manual_commission is not null then t.manual_commission
        else t.commission_target
        end as network_commission,
    CASE
        WHEN t.state = 'VALIDATED' AND r.uuid IS NOT NULL THEN r.state
        ELSE t.state
        END as status,
    t.cashback,

    r.uuid as reward_uuid,
    r.isin,
    r.current_reward_source,
    r.current_reward_target,
    r.current_reward_user,
    r.initial_price,
    r.title,
    r.end_date,
    r.created_at as start_date,
    r.origin,
    r.stopped_at
FROM transaction t
         LEFT JOIN reward r ON t.uuid = r.transaction_uuid
         left join store_visit sv ON t.store_visit_uuid = sv.uuid
         left join store s on t.store_uuid = s.uuid
where t.deleted = false;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
