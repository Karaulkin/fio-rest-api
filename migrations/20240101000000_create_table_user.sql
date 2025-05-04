-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name varchar(24) NOT NULL,
                       surname varchar(24) NOT NULL,
                       patronymic varchar(24) NOT NULL,
                       age INTEGER,
                       gender varchar(5),
                       nationality TEXT
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE users;