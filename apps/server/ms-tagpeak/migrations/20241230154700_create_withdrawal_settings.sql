-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

INSERT INTO payment_method(name, code, created_at, created_by)
VALUES ('Bank', 'bank', now(), 'system');

/*INSERT INTO payment_method(name, code, created_at, created_by)
VALUES ('PayPal', 'paypal', now(), 'system');*/

CREATE TABLE file
(
    uuid      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name      TEXT,
    extension TEXT,
    created_at          TIMESTAMP(6),
    created_by          TEXT,
    updated_at          TIMESTAMP(6),
    updated_by          TEXT,
    deleted             BOOLEAN          DEFAULT FALSE,
    deleted_at          TIMESTAMP(6),
    deleted_by          TEXT
);

CREATE TABLE user_payment_method
(
    uuid                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user"              TEXT,
    payment_method_uuid UUID,
    information         jsonb,
    file_uuid           UUID,
    created_at          TIMESTAMP(6),
    created_by          TEXT,
    updated_at          TIMESTAMP(6),
    updated_by          TEXT,
    deleted             BOOLEAN          DEFAULT FALSE,
    deleted_at          TIMESTAMP(6),
    deleted_by          TEXT,

    FOREIGN KEY (payment_method_uuid) REFERENCES payment_method (uuid),
    FOREIGN KEY (file_uuid) REFERENCES file (uuid)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
