-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO users (name, surname, patronymic, age, gender, nationality)
VALUES ('Иван', 'Иванов', 'Иванович', 30, 'M', 'Russian');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM users WHERE id = 1;