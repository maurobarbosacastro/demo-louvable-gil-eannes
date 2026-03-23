-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

DROP VIEW public_stores;
DROP VIEW cashback;

CREATE INDEX idx_validated ON transaction(updated_at)
    WHERE state = 'VALIDATED';
CREATE INDEX idx_tracked ON transaction(updated_at)
    WHERE state = 'TRACKED';
CREATE INDEX idx_expired ON reward(updated_at)
    WHERE state = 'EXPIRED';
CREATE INDEX idx_stopped ON reward(updated_at)
    WHERE state = 'STOPPED';

CREATE OR REPLACE VIEW cashback AS
SELECT t.uuid       AS transaction_uuid,
       CASE
           WHEN r.uuid IS NOT NULL THEN r."user"
           ELSE t."user"
           END      AS "user",
       sv.uuid      AS store_visit_uuid,
       s.uuid       AS store_uuid,
       s.name       AS store_name,
       s.logo       AS store_logo,
       t.order_date AS date,
       t.amount_source,
       t.amount_target,
       t.amount_user,
       t.currency_source,
       t.currency_target,
       t.created_at AS transaction_created_at,
       CASE
           WHEN t.manual_commission IS NOT NULL THEN t.manual_commission
           ELSE t.commission_target
END      AS network_commission,
       CASE
           WHEN t.state::text = 'VALIDATED'::text AND r.uuid IS NOT NULL THEN r.state
           ELSE t.state
END      AS status,
       t.cashback,
       CASE
           WHEN t.state::text = 'STOPPED' AND CURRENT_TIMESTAMP >= t.updated_at + INTERVAL '1 day' THEN true
           WHEN t.state::text = 'TRACKED' AND CURRENT_TIMESTAMP >= t.updated_at + INTERVAL '30 days' THEN true
           WHEN t.state::text = 'VALIDATED' AND CURRENT_TIMESTAMP >= t.updated_at + INTERVAL '3 days' THEN true
           WHEN t.state::text = 'EXPIRED' AND CURRENT_TIMESTAMP >= t.updated_at + INTERVAL '2 days' THEN true
           ELSE false
END as warning,
       r.uuid       AS reward_uuid,
       r.isin,
       r.current_reward_source,
       r.current_reward_target,
       r.current_reward_user,
       r.initial_price,
       r.title,
       r.end_date,
       r.created_at AS start_date,
       r.origin
FROM transaction t
         LEFT JOIN reward r ON t.uuid = r.transaction_uuid
         LEFT JOIN store_visit sv ON t.store_visit_uuid = sv.uuid
         LEFT JOIN store s ON t.store_uuid = s.uuid
WHERE t.deleted = false;

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
                    FROM cashback
                    GROUP BY store_uuid) cashback_stats ON cashback_stats.store_uuid = s.uuid
WHERE LOWER(s.state) = LOWER('ACTIVE')
  AND s.deleted = false
GROUP BY s.uuid, s.name, s.logo, s.state, visit_stats.total_visits, cashback_stats.total_cashbacks;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
