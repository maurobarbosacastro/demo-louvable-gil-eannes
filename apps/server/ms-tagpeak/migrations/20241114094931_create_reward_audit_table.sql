-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';


CREATE TABLE IF NOT EXISTS reward_audit
(
    uuid                        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    reward_uuid                 UUID REFERENCES reward (uuid),
    action                      VARCHAR(20) CHECK (action IN ('ADD', 'UPDATE', 'DELETE')),
    "user"                      TEXT,
    transaction_uuid            TEXT,
    isin                        TEXT,
    initial_reward              DOUBLE PRECISION,
    current_reward_source       DOUBLE PRECISION,
    current_reward_target       DOUBLE PRECISION,
    current_reward_user         DOUBLE PRECISION,
    currency_exchange_rate_uuid TEXT,
    currency_source             text,
    currency_target             text,
    currency_user               text,
    state                       VARCHAR(20) CHECK (state IN ('LIVE', 'STOPPED', 'FINISHED', 'REQUESTED', 'PAID')),
    initial_price               DOUBLE PRECISION,
    end_date                    DATE,
    asset_units                 INT,
    type                        VARCHAR(20) CHECK (type IN ('INVESTMENT', 'FIXED')),
    deleted                     BOOLEAN,
    created_at                  TIMESTAMP(6),
    created_by                  text,
    updated_at                  TIMESTAMP(6),
    updated_by                  text,
    deleted_at                  TIMESTAMP(6),
    deleted_by                  text
);

SELECT 'down SQL query';
-- +goose StatementEnd
