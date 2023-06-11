-- name: CreateAuthor :one
INSERT INTO authors (name, bio)
VALUES ($1, $2)
RETURNING *;

-- name: GetAuthor :one
SELECT *
FROM authors
WHERE id = $1
LIMIT 1;

-- name: UpdateAuthor :one
UPDATE authors
SET name = $2,
    bio  = $3
WHERE id = $1
RETURNING *;

-- name: PartialUpdateAuthor :one
UPDATE authors
SET name = CASE WHEN @update_name::boolean THEN @name::VARCHAR(32) ELSE name END,
    bio  = CASE WHEN @update_bio::boolean THEN @bio::TEXT ELSE bio END
WHERE id = @id
RETURNING *;

-- name: DeleteAuthor :exec
DELETE
FROM authors
WHERE id = $1;

-- name: ListAuthors :many
SELECT *
FROM authors
ORDER BY name;

-- name: TruncateAuthor :exec
TRUNCATE authors;

-- name: CreateBook :one
INSERT INTO books (title, description)
VALUES ($1, $2)
RETURNING *;

-- name: GetBook :one
SELECT *
FROM books
WHERE id = $1
LIMIT 1;

-- name: UpdateBook :one
UPDATE books
SET title = $2,
    description  = $3
WHERE id = $1
RETURNING *;

-- name: PartialUpdateBook :one
UPDATE books
SET title = CASE WHEN @update_title::boolean THEN @title::VARCHAR(32) ELSE title END,
    description  = CASE WHEN @update_description::boolean THEN @description::TEXT ELSE description END
WHERE id = @id
RETURNING *;

-- name: DeleteBook :exec
DELETE
FROM books
WHERE id = $1;

-- name: ListBooks :many
SELECT *
FROM books
ORDER BY title;

-- name: TruncateBooks :exec
TRUNCATE books;
