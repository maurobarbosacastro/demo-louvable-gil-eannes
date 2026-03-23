-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

alter table shopify_shop
add store_uuid uuid
    constraint shopify_shop_store_uuid_fk
        references store;


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
