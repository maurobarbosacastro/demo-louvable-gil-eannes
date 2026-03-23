-- +goose Up
-- +goose StatementBegin
INSERT INTO topic (name, created_at, created_by) VALUES
('all', CURRENT_TIMESTAMP, 'system'),
('base', CURRENT_TIMESTAMP, 'system'),
('silver', CURRENT_TIMESTAMP, 'system'),
('gold', CURRENT_TIMESTAMP, 'system');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM topic WHERE name IN ('all', 'base', 'silver', 'gold') AND created_by = 'system';
-- +goose StatementEnd
