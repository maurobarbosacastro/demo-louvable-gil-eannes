-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

alter table reward_audit
    add legacy_id text;


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
                              created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, deleted,
                              legacy_id)
    VALUES (NEW.uuid, trigger_action, NEW."user", NEW.transaction_uuid, NEW.isin, NEW.initial_reward,
            NEW.current_reward_source, NEW.current_reward_target, NEW.current_reward_user,
            NEW.currency_exchange_rate_uuid, NEW.currency_source, NEW.currency_target, NEW.currency_user, NEW.state,
            NEW.initial_price, NEW.end_date, NEW.asset_units, NEW.type, NEW.created_at, NEW.created_by, NEW.updated_at,
            NEW.updated_by, NEW.deleted_at, NEW.deleted_by, NEW.deleted, NEW.legacy_id);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

SELECT 'down SQL query';
-- +goose StatementEnd
