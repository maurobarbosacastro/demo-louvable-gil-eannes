-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd


CREATE TABLE IF NOT EXISTS currency_exchange_rate
(
    uuid       UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    base       text,
    rates      text,
    deleted    boolean          DEFAULT false,
    created_at TIMESTAMP(6),
    created_by text,
    updated_at TIMESTAMP(6),
    updated_by text,
    deleted_at TIMESTAMP(6),
    deleted_by text
);



-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
