-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TYPE "shop_states" AS ENUM (
    'PENDING',
    'APPROVED',
    'REJECTED',
    'SUSPENDED',
    'CLOSED'
    );

CREATE TABLE "shop" (
    "uuid" UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
    "url" TEXT NOT NULL UNIQUE,
    "state" SHOP_STATES NOT NULL DEFAULT 'PENDING',
    "access_token" TEXT,
    "installation_done" BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP(6),
    created_by TEXT,
    updated_at TIMESTAMP(6),
    updated_by TEXT,
    deleted    BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMP(6),
    deleted_by TEXT,
    PRIMARY KEY("uuid")
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
