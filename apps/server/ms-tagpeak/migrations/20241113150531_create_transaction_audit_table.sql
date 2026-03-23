-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS transaction_audit
(
    uuid              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transaction       UUID REFERENCES transaction (uuid), -- Foreign key reference to `transactions` table
    action            VARCHAR(20) CHECK (action IN ('ADD', 'UPDATE', 'DELETE')),
    "user"            text,                               -- Assuming UUID for user (could reference a `users` table if defined)
    store             UUID,                               -- Assuming UUID for store (could reference a `stores` table if defined)
    amount_source     DOUBLE PRECISION,
    amount_target     DOUBLE PRECISION,
    amount_user       DOUBLE PRECISION,
    rate              UUID,                               -- Assuming `currency_exchange_rate` is a double
    currency_source   text,                               -- Commonly, currency codes use 3 characters (e.g., USD)
    currency_target   text,
    state             VARCHAR(20) CHECK (state IN ('TRACKED', 'VALIDATED', 'REJECTED')),
    commission_source DOUBLE PRECISION,
    commission_target DOUBLE PRECISION,
    commission_user   DOUBLE PRECISION,
    order_date        DATE,
    deleted           BOOLEAN,
    store_visit       UUID,                               -- Assuming this is a foreign key (could reference a `store_visits` table if defined)
    created_at        TIMESTAMP(6),
    created_by        text,
    updated_at        TIMESTAMP(6),
    updated_by        text,
    deleted_at        TIMESTAMP(6),
    deleted_by        text
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
