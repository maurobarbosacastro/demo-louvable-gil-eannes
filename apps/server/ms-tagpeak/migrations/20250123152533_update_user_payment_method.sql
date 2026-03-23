-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

alter table user_payment_method
add state VARCHAR(20) CHECK (state IN ('PENDING', 'VALIDATED', 'REJECTED'));


insert into configuration (code, name, "value", editable, data_type, created_at, created_by, updated_at,updated_by, deleted, deleted_at, deleted_by)
values ('user_payment_iban_valid_time', 'Time, in days, that a valid user payment has the iban saved.', '30', true, 'number', now(), 'system', null, null, false, null,null);


-- +goose StatementBegin
CREATE OR REPLACE FUNCTION clear_old_file_uuids(
    p_days_interval INT,
    p_replacement_uuid UUID,
    p_update_by TEXT
)
    RETURNS TABLE (
        uuid uuid,
        file_uuid UUID,
        original_file_uuid UUID
    ) AS $$
BEGIN
    RETURN QUERY
        WITH updated_rows AS (
            UPDATE user_payment_method
                SET file_uuid = p_replacement_uuid, updated_at = now(), updated_by = p_update_by
                WHERE user_payment_method.state = 'VALIDATED'
                    AND user_payment_method.updated_at < NOW() - (p_days_interval || ' days')::INTERVAL
                    AND user_payment_method.file_uuid <> p_replacement_uuid
                RETURNING user_payment_method.uuid, user_payment_method.file_uuid as new_uuid,
                    (SELECT upm.file_uuid FROM user_payment_method as upm WHERE upm.uuid = user_payment_method.uuid) as original_file_uuid
        )
        SELECT * FROM updated_rows;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
