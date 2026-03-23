-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

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
