-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd


CREATE TABLE "shopify_shop" (
    "uuid"      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "shop_uuid" UUID,
    "user_uuid" UUID
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
