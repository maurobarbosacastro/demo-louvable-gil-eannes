-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE OR REPLACE FUNCTION action_after_reward_audit()
    RETURNS TRIGGER AS
$$
DECLARE
    trigger_action TEXT;
BEGIN
    trigger_action := TG_ARGV[0];

    INSERT INTO reward_audit (reward_uuid, action, "user", transaction_uuid, isin, initial_reward,
                              current_reward_source,
                              current_reward_target, current_reward_user, currency_exchange_rate_uuid, currency_source,
                              currency_target, currency_user, state, initial_price, end_date, asset_units, type,
                              created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, deleted)
    VALUES (NEW.uuid, trigger_action, NEW."user", NEW.transaction_uuid, NEW.isin, NEW.initial_reward,
            NEW.current_reward_source, NEW.current_reward_target, NEW.current_reward_user,
            NEW.currency_exchange_rate_uuid, NEW.currency_source, NEW.currency_target, NEW.currency_user, NEW.state,
            NEW.initial_price, NEW.end_date, NEW.asset_units, NEW.type, NEW.created_at, NEW.created_by, NEW.updated_at,
            NEW.updated_by, NEW.deleted_at, NEW.deleted_by, NEW.deleted);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER after_update_reward
    AFTER UPDATE
    ON reward
    FOR EACH ROW
    WHEN (
        NEW.deleted IS NOT DISTINCT FROM OLD.deleted AND
        OLD.deleted_at IS NOT DISTINCT FROM NEW.deleted_at AND
        OLD.deleted_by IS NOT DISTINCT FROM NEW.deleted_by
        )
EXECUTE FUNCTION action_after_reward_audit('UPDATE');


CREATE TRIGGER after_insert_reward
    AFTER INSERT
    ON reward
    FOR EACH ROW
EXECUTE FUNCTION action_after_reward_audit('ADD');

CREATE TRIGGER after_delete_reward
    AFTER UPDATE
    ON reward
    FOR EACH ROW
    WHEN ( NEW.deleted_at IS DISTINCT FROM OLD.deleted_at )
EXECUTE FUNCTION action_after_reward_audit('DELETE');

SELECT 'down SQL query';
-- +goose StatementEnd
