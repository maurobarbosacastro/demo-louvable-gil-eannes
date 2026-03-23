-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS referral
(
    uuid          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    referral_uuid UUID,
    invitee_uuid  UUID,
    created_at    TIMESTAMP(6),
    created_by    TEXT,
    updated_at    TIMESTAMP(6),
    updated_by    TEXT,
    deleted       BOOLEAN DEFAULT FALSE,
    deleted_at    TIMESTAMP(6),
    deleted_by    TEXT
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
