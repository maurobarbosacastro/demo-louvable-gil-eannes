-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

insert into partner (name, code, deep_link, deep_link_identifier, sub_identifier, created_at, created_by)
values ('CJ PT', 'cj_pt', null, '&deeplink=', '?sid=', now(), 'system'),
       ('Awin PT', 'awin_pt', '#', null, '&clickref=', now(), 'system'),
       ('Awin UK', 'awin_uk', '€', null, '&clickref=', now(), 'system'),
       ('CJ UK', 'cj_uk', null, null, '?sid=', now(), 'system'),
       ('Booking.com', 'booking', '#', null, '&label=', now(), 'system'),
       ('Rakuten', 'rakuten', null, '&murl=', '&subid=', now(), 'system'),
       ('Shopify', 'shopify', 'Shopify', 'Shopify', 'Shopify', now(), 'system'),
       ('Impact', 'impact', null, '&u=', '?subId1=', now(), 'system');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
