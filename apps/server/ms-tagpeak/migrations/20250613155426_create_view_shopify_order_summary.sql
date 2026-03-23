-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';


create or replace view shopify_order_summary as
select s.uuid as store_uuid, s.name as store_name, ss.shop_uuid as shop_uuid, count(t) as total_orders, sum(t.amount_source) as total_amount
from store s
         left join transaction t on s.uuid = t.store_uuid
         left join partner p on s.affiliate_partner_code = p.code
         left join shopify_shop ss on s.uuid = ss.store_uuid
where p.code = 'shopify'
group by s.uuid, s.name, ss.shop_uuid;



SELECT 'down SQL query';
-- +goose StatementEnd
