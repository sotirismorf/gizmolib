CREATE TABLE authors
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    bio  TEXT NOT NULL
);

CREATE TABLE books
(
    id               BIGSERIAL PRIMARY KEY,
    title            VARCHAR(64) NOT NULL,
    author_id        BIGSERIAL   NOT NULL,
    description      TEXT        NOT NULL,
	year_published   SMALLINT    NOT NULL,
	copies_total     INT         NOT NULL DEFAULT 1,
	copies_available INT         NOT NULL DEFAULT 1,
    CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES authors(id)
);

CREATE TABLE users
(
    id          BIGSERIAL PRIMARY KEY,
    username    VARCHAR(64) NOT NULL,
    password    VARCHAR(64) NOT NULL
);
