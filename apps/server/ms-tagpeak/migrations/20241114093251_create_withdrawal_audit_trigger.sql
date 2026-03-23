-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE
    OR REPLACE FUNCTION action_after_withdrawal_audit()
    RETURNS TRIGGER AS
$$
DECLARE
    trigger_action TEXT;
BEGIN
    trigger_action
        := TG_ARGV[0];

    INSERT INTO withdrawal_audit (withdrawal_id, action, "user", method, amount, details, state, completion_date, rate,
                                  currency_source, currency_target, created_at, created_by, updated_at, updated_by,
                                  deleted_at, deleted_by, deleted)
    VALUES (NEW.uuid, trigger_action, NEW."user", NEW.method, NEW.amount, NEW.details, NEW.state, NEW.completion_date,
            NEW.rate, NEW.currency_source, NEW.currency_target, NEW.created_at, NEW.created_by, NEW.updated_at,
            NEW.updated_by, NEW.deleted_at, NEW.deleted_by, NEW.deleted);

    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;



CREATE TRIGGER after_update_withdrawal
    AFTER UPDATE
    ON withdrawal
    FOR EACH ROW
    WHEN (
        NEW.deleted IS NOT DISTINCT FROM OLD.deleted AND
        OLD.deleted_at IS NOT DISTINCT FROM NEW.deleted_at AND
        OLD.deleted_by IS NOT DISTINCT FROM NEW.deleted_by
        )
EXECUTE FUNCTION action_after_withdrawal_audit('UPDATE');

CREATE TRIGGER after_insert_withdrawal
    AFTER INSERT
    ON withdrawal
    FOR EACH ROW
EXECUTE FUNCTION action_after_withdrawal_audit('ADD');

CREATE TRIGGER after_delete_withdrawal
    AFTER UPDATE
    ON withdrawal
    FOR EACH ROW
    WHEN ( NEW.deleted_at IS DISTINCT FROM OLD.deleted_at )
EXECUTE FUNCTION action_after_withdrawal_audit('DELETE');


SELECT 'down SQL query';
-- +goose StatementEnd
