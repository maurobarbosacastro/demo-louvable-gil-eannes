-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- ADD CONSTRAINTS AND CHANGE COLUMN TYPES
ALTER TABLE language
    ADD CONSTRAINT unique_code UNIQUE (code);

alter table store
    ALTER COLUMN average_reward_activation_time TYPE TEXT USING average_reward_activation_time::TEXT;


-- DROP FKeys
alter table store
    drop constraint store_category_uuid_fkey;

alter table store
    drop constraint store_language_uuid_fkey;

alter table store
    drop constraint store_affiliate_partner_uuid_fkey;


-- RENAME COLUMNS
alter table store
    rename column language_uuid to language_code;

alter table store
    alter column language_code type TEXT using language_code::TEXT;

alter table store
    rename column affiliate_partner_uuid to affiliate_partner_code;

alter table store
    alter column affiliate_partner_code type TEXT using affiliate_partner_code::TEXT;


 -- ADD FKeys
alter table store
    add constraint store_language_code_fkey
        foreign key (language_code) references language (code);

alter table store
    add constraint store_affiliate_partner_code_fkey
        foreign key (affiliate_partner_code) references partner (code);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
