-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd


-- Withdrawal
ALTER TABLE withdrawal
drop column rate;

ALTER TABLE withdrawal
add column currency_exchange_rate_uuid UUID
constraint withdrawal_currency_exchange_rate_uuid_fk REFERENCES currency_exchange_rate (uuid);

alter table withdrawal
rename column amount to amount_source;

ALTER table withdrawal
add column amount_target double precision;

-- Withdrawal Audit
ALTER TABLE withdrawal_audit
drop column rate;

ALTER TABLE withdrawal_audit
add column currency_exchange_rate_uuid UUID;

alter table withdrawal_audit
rename column amount to amount_source;

ALTER table withdrawal_audit
add column amount_target double precision;



-- +goose StatementBegin
-- Triggers
CREATE OR REPLACE FUNCTION action_after_withdrawal_audit()
    RETURNS TRIGGER AS
$$
DECLARE
    trigger_action TEXT;
BEGIN
    trigger_action := TG_ARGV[0];

    INSERT INTO withdrawal_audit (withdrawal_id, action, "user", method, amount_source, details, state, completion_date,
                                  currency_exchange_rate_uuid, amount_target, currency_source, currency_target, created_at,
                                  created_by, updated_at, updated_by, deleted_at, deleted_by, deleted)
    VALUES (NEW.uuid, trigger_action, NEW."user", NEW.method, NEW.amount_source, NEW.details, NEW.state, NEW.completion_date,
            NEW.currency_exchange_rate_uuid, NEW.amount_target, NEW.currency_source, NEW.currency_target, NEW.created_at,
            NEW.created_by, NEW.updated_at, NEW.updated_by, NEW.deleted_at, NEW.deleted_by, NEW.deleted);

    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;
-- +goose StatementEnd




-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
