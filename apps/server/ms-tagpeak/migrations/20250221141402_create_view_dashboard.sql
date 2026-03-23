-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE OR REPLACE VIEW dashboard AS
SELECT
    EXTRACT(MONTH FROM t.order_date) AS month,
    EXTRACT(YEAR FROM t.order_date) AS year,
    COALESCE(COUNT(DISTINCT t."user"), 0) AS active_users,
    COALESCE(COUNT(DISTINCT t.uuid), 0) AS num_transactions,
    COALESCE(SUM(t.amount_target), 0) AS total_gmv,
    COALESCE(AVG(t.amount_target), 0) AS avg_transaction_amount,
    COALESCE(SUM(
                     CASE
                         WHEN t.manual_commission IS NOT NULL THEN t.manual_commission
                         ELSE t.commission_target
                         END
             ), 0) AS total_revenue
FROM transaction t
WHERE t.deleted = false and t.deleted_at IS NULL
GROUP BY year, month
ORDER BY year DESC, month DESC;


CREATE OR REPLACE VIEW cashback_dashboard AS
SELECT
    status,
    SUM(value) AS value,
    SUM(count) AS count,
    SUM(warning) AS warning
FROM (
        SELECT
            CASE
                WHEN t.state = 'VALIDATED' AND r.uuid IS NOT NULL THEN r.state
                ELSE t.state
            END AS status,
            CASE
                WHEN t.state = 'STOPPED' AND CURRENT_TIMESTAMP >= t.updated_at + INTERVAL '1 day' THEN 1
                WHEN t.state = 'TRACKED' AND CURRENT_TIMESTAMP >= t.updated_at + INTERVAL '30 days' THEN 1
                WHEN t.state = 'VALIDATED' AND CURRENT_TIMESTAMP >= t.updated_at + INTERVAL '3 days' THEN 1
                WHEN t.state = 'EXPIRED' AND CURRENT_TIMESTAMP >= t.updated_at + INTERVAL '2 days' THEN 1
                ELSE 0
            END AS warning,
            CASE
                WHEN t.state IN ('TRACKED', 'VALIDATED') THEN t.cashback
                ELSE COALESCE(r.current_reward_target, 0)
            END AS value,
            1 AS count
        FROM transaction t
            LEFT JOIN reward r ON t.uuid = r.transaction_uuid WHERE t.deleted = false and t.deleted_at IS NULL
    ) sub
GROUP BY status;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
