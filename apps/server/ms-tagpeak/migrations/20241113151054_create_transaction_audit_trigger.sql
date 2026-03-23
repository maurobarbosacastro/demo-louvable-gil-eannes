-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';


-- Define the trigger function
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
                                   updated_by, deleted_at, deleted_by, deleted)
    VALUES (NEW.uuid, trigger_action, NEW."user", NEW.store_uuid, NEW.amount_source, NEW.amount_target, NEW.amount_user,
            NEW.currency_exchange_rate_uuid, NEW.currency_source, NEW.currency_target, NEW.state, NEW.commission_source,
            NEW.commission_target, NEW.commission_user, NEW.order_date, NEW.store_visit_uuid, NEW.created_at,
            NEW.created_by, NEW.updated_at, NEW.updated_by, NEW.deleted_at, NEW.deleted_by, NEW.deleted);

-- Return the modified row (necessary for AFTER UPDATE and INSERT triggers)
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;


-- Create the AFTER UPDATE trigger
CREATE TRIGGER after_update_transaction
    AFTER UPDATE
    ON transaction
    FOR EACH ROW
    WHEN (
        NEW.deleted IS NOT DISTINCT FROM OLD.deleted AND
        OLD.deleted_at IS NOT DISTINCT FROM NEW.deleted_at AND
        OLD.deleted_by IS NOT DISTINCT FROM NEW.deleted_by
        )
EXECUTE FUNCTION action_after_transaction_audit('UPDATE');

-- Create the AFTER INSERT trigger
CREATE TRIGGER after_create_transaction
    AFTER INSERT
    ON transaction
    FOR EACH ROW
EXECUTE FUNCTION action_after_transaction_audit('ADD');

-- Create the AFTER DELETE trigger
CREATE TRIGGER after_delete_transaction
    BEFORE UPDATE
    ON transaction
    FOR EACH ROW
    WHEN ( NEW.deleted_at IS DISTINCT FROM OLD.deleted_at )
EXECUTE FUNCTION action_after_transaction_audit('DELETE');

SELECT 'down SQL query';
-- +goose StatementEnd
