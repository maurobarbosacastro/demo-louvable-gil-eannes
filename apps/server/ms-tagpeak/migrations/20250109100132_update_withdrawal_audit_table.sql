-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

alter table withdrawal_audit
    add legacy_id text;


CREATE OR REPLACE FUNCTION action_after_withdrawal_audit()
    RETURNS TRIGGER AS
$$
DECLARE
    trigger_action TEXT;
BEGIN
    trigger_action := TG_ARGV[0];

    INSERT INTO withdrawal_audit (withdrawal_id, action, "user", user_method, amount_source, details, state, completion_date,
                                  currency_exchange_rate_uuid, amount_target, currency_source, currency_target, created_at,
                                  created_by, updated_at, updated_by, deleted_at, deleted_by, deleted, legacy_id)
    VALUES (NEW.uuid, trigger_action, NEW."user", NEW.user_method, NEW.amount_source, NEW.details, NEW.state, NEW.completion_date,
            NEW.currency_exchange_rate_uuid, NEW.amount_target, NEW.currency_source, NEW.currency_target, NEW.created_at,
            NEW.created_by, NEW.updated_at, NEW.updated_by, NEW.deleted_at, NEW.deleted_by, NEW.deleted, NEW.legacy_id);

    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

SELECT 'down SQL query';
-- +goose StatementEnd
