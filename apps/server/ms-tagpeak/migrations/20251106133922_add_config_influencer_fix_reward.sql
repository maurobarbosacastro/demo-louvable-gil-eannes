-- +goose Up
-- +goose StatementBegin
INSERT INTO configuration (code, name, value, editable, data_type, created_at, created_by)
VALUES ('influencer_default_amount', 'Default amount influencers will received upon a validated transaction', '5', true, 'number', now(), 'system');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM configuration WHERE key = 'influencer_default_amount';
-- +goose StatementEnd
