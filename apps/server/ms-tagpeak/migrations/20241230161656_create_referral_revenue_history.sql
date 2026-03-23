-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS referral_revenue_history (
    uuid UUID                   PRIMARY KEY DEFAULT uuid_generate_v4(),
    referral_uuid               UUID,
    transaction_uuid            UUID NULL,
    reward_uuid                 UUID NULL,
    amount                      DOUBLE PRECISION,

    created_at                  TIMESTAMP(6),
    created_by                  TEXT,
    updated_at                  TIMESTAMP(6),
    updated_by                  TEXT,
    deleted_at                  TIMESTAMP(6),
    deleted_by                  TEXT,


    FOREIGN KEY (referral_uuid) REFERENCES referral (uuid),
    FOREIGN KEY (transaction_uuid) REFERENCES transaction (uuid),
    FOREIGN KEY (reward_uuid) REFERENCES reward (uuid)

    );

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
