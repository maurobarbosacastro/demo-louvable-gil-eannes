-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE OR REPLACE VIEW currency_dashboard AS
SELECT r.currency_user as currency,
       r.state as state,
       COUNT(*) as total_rewards
FROM reward as r
GROUP BY r.currency_user, r.state;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
