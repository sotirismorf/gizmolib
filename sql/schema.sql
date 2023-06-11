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
    description TEXT NOT NULL
);

-- CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES authors(id)
