-- +goose Up
-- +goose StatementBegin

-- Create new enum for join table states
CREATE TYPE join_table_notification_state AS ENUM ('pending', 'delivered', 'canceled', 'error');

-- Remove existing default before type change
ALTER TABLE notification ALTER COLUMN state DROP DEFAULT;

-- Update notification state enum to new values
ALTER TYPE notification_state RENAME TO notification_state_old;
CREATE TYPE notification_state AS ENUM ('draft', 'scheduled', 'processed', 'error', 'cancelled');

-- Update existing notification records to map old states to new states
-- pending -> draft, delivered -> processed, canceled -> cancelled, error -> error
ALTER TABLE notification ALTER COLUMN state TYPE notification_state USING
  CASE
    WHEN state::text = 'pending' THEN 'draft'::notification_state
    WHEN state::text = 'delivered' THEN 'processed'::notification_state
    WHEN state::text = 'canceled' THEN 'cancelled'::notification_state
    WHEN state::text = 'error' THEN 'error'::notification_state
    ELSE 'draft'::notification_state
  END;

-- Drop old enum type
DROP TYPE notification_state_old;

-- Set new default value for notification state
ALTER TABLE notification ALTER COLUMN state SET DEFAULT 'draft'::notification_state;

-- Add state column to notification_user_tokens join table
ALTER TABLE notification_user_tokens
ADD COLUMN state join_table_notification_state NOT NULL DEFAULT 'pending';

-- Add state column to notification_topics join table
ALTER TABLE notification_topics
ADD COLUMN state join_table_notification_state NOT NULL DEFAULT 'pending';

-- Add indexes for the new state columns
CREATE INDEX idx_notification_user_tokens_state ON notification_user_tokens(state);
CREATE INDEX idx_notification_topics_state ON notification_topics(state);

-- Add created_at and updated_at timestamps to join tables for audit trail
ALTER TABLE notification_user_tokens
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP;

ALTER TABLE notification_topics
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Remove added columns from join tables
ALTER TABLE notification_user_tokens
DROP COLUMN IF EXISTS state,
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE notification_topics
DROP COLUMN IF EXISTS state,
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;

-- Remove default before type change
ALTER TABLE notification ALTER COLUMN state DROP DEFAULT;

-- Restore original notification state enum
ALTER TYPE notification_state RENAME TO notification_state_new;
CREATE TYPE notification_state AS ENUM ('pending', 'delivered', 'canceled', 'error');

-- Update existing notification records back to old states
ALTER TABLE notification ALTER COLUMN state TYPE notification_state USING
  CASE
    WHEN state::text = 'draft' THEN 'pending'::notification_state
    WHEN state::text = 'scheduled' THEN 'pending'::notification_state
    WHEN state::text = 'processed' THEN 'delivered'::notification_state
    WHEN state::text = 'error' THEN 'error'::notification_state
    WHEN state::text = 'cancelled' THEN 'canceled'::notification_state
    ELSE 'pending'::notification_state
  END;

-- Drop new enum type
DROP TYPE notification_state_new;

-- Restore original default
ALTER TABLE notification ALTER COLUMN state SET DEFAULT 'pending'::notification_state;

-- Drop join table state enum
DROP TYPE IF EXISTS join_table_notification_state;

-- Drop indexes
DROP INDEX IF EXISTS idx_notification_user_tokens_state;
DROP INDEX IF EXISTS idx_notification_topics_state;

-- +goose StatementEnd