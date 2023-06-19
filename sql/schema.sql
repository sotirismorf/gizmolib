CREATE TABLE authors
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    bio  TEXT NOT NULL
);

CREATE TABLE books
(
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(64) NOT NULL,
    author_id   BIGSERIAL NOT NULL,
    description TEXT NOT NULL,
    CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES authors(id)
);

CREATE TABLE users
(
    id          BIGSERIAL PRIMARY KEY,
    username    VARCHAR(64) NOT NULL,
    password    VARCHAR(64) NOT NULL
);

INSERT INTO users (username, password)
VALUES ('username1', '$2a$10$gcF3GsYfgk5VPS4cKMq/0e318zQWquObmy1wtjUi2jifK2mKtnxyi')
RETURNING *;
