-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

alter table transaction
add cashback double precision;

alter table reward
alter column asset_units type double precision using asset_units::double precision;


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
