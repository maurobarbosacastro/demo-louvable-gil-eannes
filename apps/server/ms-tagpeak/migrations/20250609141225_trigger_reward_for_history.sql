-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- 1. Create the trigger function
CREATE
    OR REPLACE FUNCTION trg_insert_reward_history()
    RETURNS TRIGGER AS
$$
BEGIN
    -- Log the new reward UUID being inserted into history
    RAISE NOTICE 'Inserting reward_history for reward_uuid: %', NEW.uuid;

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
        NEW.uuid,
        NEW.initial_price,
        NEW.asset_units,
        NEW.current_reward_user,
        NEW.deleted,
        NEW.created_at,
        NEW.created_by,
        NEW.updated_at,
        NEW.updated_by,
        NEW.deleted_at,
        NEW.deleted_by,
        uuid_generate_v4()
    );

    RAISE NOTICE 'Inserted reward_history record for reward_uuid: %', NEW.uuid;

    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

-- 2. Create the trigger on the reward table
CREATE TRIGGER insert_reward_history_after_insert
    AFTER INSERT ON reward
    FOR EACH ROW
EXECUTE FUNCTION trg_insert_reward_history();

SELECT 'down SQL query';
-- +goose StatementEnd
