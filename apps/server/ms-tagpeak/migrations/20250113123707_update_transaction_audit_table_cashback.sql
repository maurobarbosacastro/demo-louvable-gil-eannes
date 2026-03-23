-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

alter table transaction_audit
    add cashback double precision;

alter table reward_audit
    alter column asset_units type double precision using asset_units::double precision;

alter table referral_revenue_history
    add deleted BOOLEAN DEFAULT FALSE;

CREATE
    OR REPLACE FUNCTION action_after_transaction_audit()
    RETURNS TRIGGER AS
$$
DECLARE
    trigger_action TEXT;

BEGIN
    trigger_action
        := TG_ARGV[0];

    -- Insert old values and the action type into transaction_audit
    INSERT INTO transaction_audit (transaction, action, "user", store, amount_source, amount_target, amount_user, rate,
                                   currency_source, currency_target, state, commission_source, commission_target,
                                   commission_user, order_date, store_visit, created_at, created_by, updated_at,
                                   updated_by, deleted_at, deleted_by, deleted, legacy_id, source_id, cashback)
    VALUES (NEW.uuid, trigger_action, NEW."user", NEW.store_uuid, NEW.amount_source, NEW.amount_target, NEW.amount_user,
            NEW.currency_exchange_rate_uuid, NEW.currency_source, NEW.currency_target, NEW.state, NEW.commission_source,
            NEW.commission_target, NEW.commission_user, NEW.order_date, NEW.store_visit_uuid, NEW.created_at,
            NEW.created_by, NEW.updated_at, NEW.updated_by, NEW.deleted_at, NEW.deleted_by, NEW.deleted, NEW.legacy_id,
            NEW.source_id, NEW.cashback);

-- Return the modified row (necessary for AFTER UPDATE and INSERT triggers)
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;



SELECT 'down SQL query';
-- +goose StatementEnd
