-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

alter table transaction_audit
    add manual_commission double precision,
    add amount_vat_source double precision,
    add amount_vat_target double precision,
    add amount_vat_user double precision,
    add description text;


alter table transaction
    add manual_commission double precision,
    add amount_vat_source double precision,
    add amount_vat_target double precision,
    add amount_vat_user double precision,
    add description text;


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
                                   updated_by, deleted_at, deleted_by, deleted, legacy_id, source_id, cashback, description, amount_vat_source, amount_vat_target, amount_vat_user)
    VALUES (NEW.uuid, trigger_action, NEW."user", NEW.store_uuid, NEW.amount_source, NEW.amount_target, NEW.amount_user,
            NEW.currency_exchange_rate_uuid, NEW.currency_source, NEW.currency_target, NEW.state, NEW.commission_source,
            NEW.commission_target, NEW.commission_user, NEW.order_date, NEW.store_visit_uuid, NEW.created_at,
            NEW.created_by, NEW.updated_at, NEW.updated_by, NEW.deleted_at, NEW.deleted_by, NEW.deleted, NEW.legacy_id,
            NEW.source_id, NEW.cashback, NEW.description, NEW.amount_vat_source, NEW.amount_vat_target, NEW.amount_vat_user);

-- Return the modified row (necessary for AFTER UPDATE and INSERT triggers)
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;



SELECT 'down SQL query';
-- +goose StatementEnd
