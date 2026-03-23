-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE
    OR REPLACE FUNCTION trg_reward_override_price_change()
    RETURNS TRIGGER AS
$$
BEGIN
    -- Check if override_price changed
    IF OLD.override_price IS DISTINCT FROM NEW.override_price THEN

        INSERT INTO reward_history (
            reward_uuid,
            rate,
            units,
            cash_reward,
            deleted,
            created_at,
            created_by,
            updated_at,
            updated_by,
            deleted_at,
            deleted_by,
            uuid
        )
        VALUES (
            NEW.uuid,                -- reward_uuid
            NEW.override_price,      -- rate mapped from override_price
            NEW.asset_units,         -- units
            NEW.current_reward_user, -- cash_reward
            NEW.deleted,
            CURRENT_TIMESTAMP,       -- created_at (use current time for history)
            NEW.updated_by,          -- created_by (use updater)
            NULL,                    -- updated_at
            NULL,                    -- updated_by
            NULL,                    -- deleted_at
            NULL,                    -- deleted_by
            uuid_generate_v4()        -- new uuid for reward_history
        );

        RAISE NOTICE 'Inserted reward_history record for override_price change on reward_uuid: %', NEW.uuid;
    END IF;

    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

CREATE TRIGGER trg_reward_override_price_update
    AFTER UPDATE ON reward
    FOR EACH ROW
EXECUTE FUNCTION trg_reward_override_price_change();

SELECT 'down SQL query';
-- +goose StatementEnd
