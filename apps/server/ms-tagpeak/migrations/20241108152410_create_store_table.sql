-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE store
(
    uuid                           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name                           TEXT NOT NULL,
    logo                           TEXT,
    short_description              TEXT,
    description                    TEXT,
    url_slug                       TEXT,
    initial_reward                 FLOAT,
    average_reward_activation_time FLOAT,
    state                          VARCHAR(50),
    keywords                       TEXT,
    affiliate_link                 TEXT,
    store_url                      TEXT,
    terms_and_conditions           TEXT,
    cashback_type                  VARCHAR(50),
    cashback_value                 FLOAT,
    percentage_cashout             FLOAT,
    meta_title                     TEXT,
    meta_keywords                  TEXT,
    meta_description               TEXT,
    created_at                     TIMESTAMP(6),
    created_by                     TEXT,
    updated_at                     TIMESTAMP(6),
    updated_by                     TEXT,
    deleted                        BOOLEAN          DEFAULT FALSE,
    deleted_at                     TIMESTAMP(6),
    deleted_by                     TEXT,
    category_uuid                  UUID, -- Foreign key to category table
    country_uuid                   UUID, -- Foreign key to country table
    language_uuid                  UUID, -- Foreign key to language table
    affiliate_partner_uuid         UUID, -- Foreign key to affiliate/partner table

    -- Foreign key constraints
    FOREIGN KEY (category_uuid) REFERENCES category (uuid),
    FOREIGN KEY (country_uuid) REFERENCES country (uuid),
    FOREIGN KEY (language_uuid) REFERENCES language (uuid),
    FOREIGN KEY (affiliate_partner_uuid) REFERENCES partner (uuid)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
