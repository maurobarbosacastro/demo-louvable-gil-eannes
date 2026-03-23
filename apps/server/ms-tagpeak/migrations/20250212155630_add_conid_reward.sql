-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

alter table reward
add conid text;

CREATE
    or replace VIEW cashback AS
SELECT t.uuid  as transaction_uuid,
    case
        when r.uuid is not null then r."user"
        else t."user"
        end as "user",
    sv.uuid as store_visit_uuid,
    s.uuid  as store_uuid,
    s.name  as store_name,
    s.logo  as store_logo,
    t.order_date as date,
    t.amount_source,
    t.amount_target,
    t.amount_user,
    t.currency_source,
    t.currency_target,
    CASE
        when t.manual_commission is not null then t.manual_commission
        else t.commission_target
        end
        as network_commission,
    CASE
        WHEN t.state = 'VALIDATED' AND r.uuid IS NOT NULL THEN r.state
        ELSE t.state
        END
        as status,
    t.cashback,

    r.uuid as reward_uuid,
    r.isin,
    r.current_reward_source,
    r.current_reward_target,
    r.current_reward_user,
    r.initial_price,
    r.title,
    r.end_date,
    r.created_at as start_date,
    r.origin,
    r.conid
FROM transaction t
         LEFT JOIN reward r ON t.uuid = r.transaction_uuid
         left join store_visit sv ON t.store_visit_uuid = sv.uuid
         left join store s on t.store_uuid = s.uuid
where t.deleted = false;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
