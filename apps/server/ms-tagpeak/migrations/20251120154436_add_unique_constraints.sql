-- +goose Up
-- +goose StatementBegin
ALTER TABLE reward
ADD CONSTRAINT reward_user_transaction_origin_unique
UNIQUE ("user", transaction_uuid, origin);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
ALTER TABLE reward
DROP CONSTRAINT reward_user_transaction_origin_unique;
-- +goose StatementEnd
