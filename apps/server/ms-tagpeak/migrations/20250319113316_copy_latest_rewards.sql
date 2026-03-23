-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE
OR REPLACE FUNCTION copy_latest_rewards(reward_uuids uuid[])
    RETURNS void AS $$
BEGIN
INSERT INTO reward_history (reward_uuid,
                            rate,
                            units,
                            cash_reward,
                            deleted,
                            created_at,
                            created_by,
                            updated_at,
                            updated_by,
                            deleted_at,
                            deleted_by)
SELECT rh.reward_uuid,
       rh.rate,
       rh.units,
       rh.cash_reward,
       rh.deleted,
       CURRENT_TIMESTAMP(6), -- new created_at timestamp
       rh.created_by,
       CURRENT_TIMESTAMP(6), -- reset updated_at
       NULL,                 -- reset updated_by
       rh.deleted_at,
       rh.deleted_by
FROM reward_history rh
         INNER JOIN (SELECT reward_uuid, MAX(created_at) as max_created_at
                     FROM reward_history
                     WHERE reward_uuid = ANY ($1)
                     GROUP BY reward_uuid) latest ON rh.reward_uuid = latest.reward_uuid
    AND rh.created_at = latest.max_created_at;
END;
$$
LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
