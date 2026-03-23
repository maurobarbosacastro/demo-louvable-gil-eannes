-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- Change APPROVED to ACTIVE in shop_states enum
ALTER TYPE shop_states RENAME VALUE 'APPROVED' TO 'ACTIVE';


-- Update default value for state column in shop table from PENDING to ACTIVE
ALTER TABLE shop
ALTER COLUMN state SET DEFAULT 'ACTIVE';

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- Revert ACTIVE back to APPROVED in shop_states enum
ALTER TYPE shop_states RENAME VALUE 'ACTIVE' TO 'APPROVED';

-- Update default value for state column in shop table from PENDING to ACTIVE
ALTER TABLE shop
ALTER COLUMN state SET DEFAULT 'PENDING';
