-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE OR REPLACE FUNCTION update_shopify_actionable_transactions_state()
    RETURNS TABLE
    (
        trx_uuid uuid,
        trx_source_id text,
        trx_amount_source double precision,
        trx_currency_source text,
        trx_commission_source double precision,
        trx_order_date date,
        trx_store_visit_uuid uuid,
        trx_user text
    )
AS
$$
BEGIN
    RETURN QUERY
        WITH updated_rows AS (
            UPDATE transaction
                SET state = 'VALIDATED', updated_at = now(), updated_by = 'system'
                WHERE uuid IN (SELECT transaction_uuid
                               FROM shopify_actionable_transactions
                               WHERE remaining_time = 0)
                RETURNING uuid, transaction.source_id, transaction.amount_source, transaction.currency_source, transaction.commission_source, transaction.order_date, transaction.store_visit_uuid, transaction."user")
        SELECT *
        FROM updated_rows;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
