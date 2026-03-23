-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS transaction
(
    uuid                        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    amount_source               DOUBLE PRECISION,
    amount_target               DOUBLE PRECISION,
    amount_user                 DOUBLE PRECISION,
    currency_source             text,
    currency_target             text,
    state                       VARCHAR(20) CHECK (state IN ('TRACKED', 'VALIDATED', 'REJECTED')),
    commission_source           DOUBLE PRECISION,
    commission_target           DOUBLE PRECISION,
    commission_user             DOUBLE PRECISION,
    order_date                  DATE,
    "user"                      text,
    deleted                     BOOLEAN          DEFAULT FALSE,
    store_uuid                  UUID,
    store_visit_uuid            UUID,
    currency_exchange_rate_uuid UUID,

    created_at                  TIMESTAMP(6),
    created_by                  TEXT,
    updated_at                  TIMESTAMP(6),
    updated_by                  TEXT,
    deleted_at                  TIMESTAMP(6),
    deleted_by                  TEXT,

    -- Foreign key constraints
    FOREIGN KEY (store_uuid) REFERENCES store (uuid),
    FOREIGN KEY (store_visit_uuid) REFERENCES store_visit (uuid),
    FOREIGN KEY (currency_exchange_rate_uuid) REFERENCES currency_exchange_rate (uuid)

);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
