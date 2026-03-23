-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE reward
DROP CONSTRAINT reward_state_check;

ALTER TABLE reward
ADD CONSTRAINT reward_state_check
    CHECK (state IN ('LIVE', 'STOPPED', 'FINISHED', 'REQUESTED', 'PAID', 'EXPIRED'));

ALTER TABLE reward_audit
DROP CONSTRAINT reward_audit_state_check;

ALTER TABLE reward_audit
ADD CONSTRAINT reward_audit_state_check
    CHECK (state IN ('LIVE', 'STOPPED', 'FINISHED', 'REQUESTED', 'PAID', 'EXPIRED'));

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
