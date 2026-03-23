-- +goose Up
-- +goose StatementBegin
ALTER TABLE reward
    DROP CONSTRAINT IF EXISTS reward_origin_check;

ALTER TABLE reward
    ADD CONSTRAINT reward_origin_check CHECK (origin IN ('PURCHASE', 'REFERRAL', 'COMMISSION'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE reward
    DROP CONSTRAINT IF EXISTS reward_origin_check;

ALTER TABLE reward
    ADD CONSTRAINT reward_origin_check CHECK (origin IN ('PURCHASE', 'REFERRAL'));
-- +goose StatementEnd
