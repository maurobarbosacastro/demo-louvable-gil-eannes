-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE store ADD COLUMN position INTEGER NULL;

DROP VIEW public_stores;
CREATE OR REPLACE VIEW public_stores
            (store_uuid, name, logo, state, created_at, position, total_visits, total_cashbacks,
             countries_code, categories_code) AS
SELECT
    s.uuid                                              AS store_uuid,
    s.name,
    s.logo,
    s.state,
    s.created_at,
    s.position,
    COALESCE(visit_stats.total_visits, 0::bigint)       AS total_visits,
    COALESCE(cashback_stats.total_cashbacks, 0::bigint) AS total_cashbacks,
    array_agg(DISTINCT country.abbreviation)            AS countries_code,
    array_agg(DISTINCT c.code)                          AS categories_code
FROM store s
         LEFT JOIN store_country
                   ON store_country.store_uuid = s.uuid
         LEFT JOIN country
                   ON country.abbreviation = store_country.country_code
         LEFT JOIN store_category
                   ON store_category.store_uuid = s.uuid
         LEFT JOIN category c
                   ON c.code = store_category.category_code
         LEFT JOIN (
    SELECT store_visit.store_uuid,
           count(*) AS total_visits
    FROM store_visit
    GROUP BY store_visit.store_uuid
) visit_stats
                   ON visit_stats.store_uuid = s.uuid
         LEFT JOIN (
    SELECT cashback.store_uuid,
           count(*) AS total_cashbacks
    FROM cashback
    GROUP BY cashback.store_uuid
) cashback_stats
                   ON cashback_stats.store_uuid = s.uuid
WHERE lower(s.state::text) = lower('ACTIVE')
  AND s.deleted = false
GROUP BY
    s.uuid, s.name, s.logo, s.state, s.created_at, s.position,
    visit_stats.total_visits, cashback_stats.total_cashbacks;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP VIEW public_stores;
CREATE OR REPLACE VIEW public_stores
            (store_uuid, name, logo, state, created_at, total_visits, total_cashbacks,
             countries_code, categories_code) AS
SELECT
    s.uuid                                              AS store_uuid,
    s.name,
    s.logo,
    s.state,
    s.created_at,
    COALESCE(visit_stats.total_visits, 0::bigint)       AS total_visits,
    COALESCE(cashback_stats.total_cashbacks, 0::bigint) AS total_cashbacks,
    array_agg(DISTINCT country.abbreviation)            AS countries_code,
    array_agg(DISTINCT c.code)                          AS categories_code
FROM store s
         LEFT JOIN store_country
                   ON store_country.store_uuid = s.uuid
         LEFT JOIN country
                   ON country.abbreviation = store_country.country_code
         LEFT JOIN store_category
                   ON store_category.store_uuid = s.uuid
         LEFT JOIN category c
                   ON c.code = store_category.category_code
         LEFT JOIN (
    SELECT store_visit.store_uuid,
           count(*) AS total_visits
    FROM store_visit
    GROUP BY store_visit.store_uuid
) visit_stats
                   ON visit_stats.store_uuid = s.uuid
         LEFT JOIN (
    SELECT cashback.store_uuid,
           count(*) AS total_cashbacks
    FROM cashback
    GROUP BY cashback.store_uuid
) cashback_stats
                   ON cashback_stats.store_uuid = s.uuid
WHERE lower(s.state::text) = lower('ACTIVE')
  AND s.deleted = false
GROUP BY
    s.uuid, s.name, s.logo, s.state, s.created_at,
    visit_stats.total_visits, cashback_stats.total_cashbacks;

ALTER TABLE store DROP COLUMN position;
