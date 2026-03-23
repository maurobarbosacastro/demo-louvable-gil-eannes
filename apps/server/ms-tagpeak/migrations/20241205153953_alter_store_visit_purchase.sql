-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE store_visit
ALTER COLUMN purchase TYPE bool
    USING CASE WHEN purchase > 0 THEN true ELSE false END;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
