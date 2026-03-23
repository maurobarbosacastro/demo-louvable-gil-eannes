-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

INSERT INTO category (name, code, created_at, created_by)
VALUES ('Fashion', 'fashion', NOW(), 'system'),
       ('Sports', 'sports', NOW(), 'system'),
       ('Luxury', 'luxury', NOW(), 'system'),
       ('Electronics', 'electronics', NOW(), 'system'),
       ('Food', 'food', NOW(), 'system'),
       ('Travel', 'travel', NOW(), 'system'),
       ('Baby', 'baby', NOW(), 'system'),
       ('Department Stores', 'department_stores', NOW(), 'system'),
       ('Personal Care', 'personal_care', NOW(), 'system'),
       ('Home', 'home', NOW(), 'system');

SELECT 'down SQL query';
-- +goose StatementEnd
