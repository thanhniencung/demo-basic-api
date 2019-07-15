-- +migrate Up
CREATE TABLE "users"
(
    "name" text NOT NULL,
    "email" text NOT NULL
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;



