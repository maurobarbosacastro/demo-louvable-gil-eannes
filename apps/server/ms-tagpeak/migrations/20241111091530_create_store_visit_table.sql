-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd


CREATE TABLE IF NOT EXISTS store_visit
(
    uuid       UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user"     VARCHAR(255),
    reference  VARCHAR(255),
    purchase   FLOAT,
    store_uuid UUID,
    created_at TIMESTAMP(6),
    created_by VARCHAR(255),
    updated_at TIMESTAMP(6),
    updated_by VARCHAR(255),
    deleted    BOOLEAN          DEFAULT FALSE,
    deleted_at TIMESTAMP(6),
    deleted_by TEXT,
    -- Foreign key constraints
    FOREIGN KEY (store_uuid) REFERENCES store (uuid)
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
