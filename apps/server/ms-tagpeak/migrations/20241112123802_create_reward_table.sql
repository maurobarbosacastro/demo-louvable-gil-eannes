-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS reward
(
    uuid                        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user"                      TEXT,
    transaction_uuid            UUID,
    isin                        TEXT,
    initial_reward              DOUBLE PRECISION,
    current_reward_source       DOUBLE PRECISION,
    current_reward_target       DOUBLE PRECISION,
    current_reward_user         DOUBLE PRECISION,
    currency_exchange_rate_uuid UUID,
    currency_source             TEXT,
    currency_target             TEXT,
    currency_user               TEXT,
    state                       VARCHAR(20) CHECK (state IN ('LIVE', 'STOPPED', 'FINISHED', 'REQUESTED', 'PAID')),
    initial_price               DOUBLE PRECISION,
    end_date                    DATE,
    asset_units                 INT,
    type                        VARCHAR(20) CHECK (type IN ('INVESTMENT', 'FIXED')),
    deleted                     BOOLEAN          DEFAULT FALSE,
    created_at                  TIMESTAMP(6),
    created_by                  TEXT,
    updated_at                  TIMESTAMP(6),
    updated_by                  TEXT,
    deleted_at                  TIMESTAMP(6),
    deleted_by                  TEXT,

    FOREIGN KEY (transaction_uuid) REFERENCES transaction (uuid),
    FOREIGN KEY (currency_exchange_rate_uuid) REFERENCES currency_exchange_rate (uuid)


);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
