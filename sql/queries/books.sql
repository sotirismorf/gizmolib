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
