CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name varchar(12) NOT NULL,
                       surname varchar(24) NOT NULL,
                       patronymic varchar(14) NOT NULL,
                       age INTEGER,
                       gender varchar(5),
                       nationality TEXT
);