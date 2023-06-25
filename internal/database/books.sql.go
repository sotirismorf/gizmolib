// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: books.sql

package database

import (
	"context"
)

const createBook = `-- name: CreateBook :one
INSERT INTO books (title, description, author_id, year_published, copies_total, copies_available)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, title, author_id, description, year_published, copies_total, copies_available
`

type CreateBookParams struct {
	Title           string
	Description     string
	AuthorID        int64
	YearPublished   int16
	CopiesTotal     int32
	CopiesAvailable int32
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.Title,
		arg.Description,
		arg.AuthorID,
		arg.YearPublished,
		arg.CopiesTotal,
		arg.CopiesAvailable,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.AuthorID,
		&i.Description,
		&i.YearPublished,
		&i.CopiesTotal,
		&i.CopiesAvailable,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE
FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBook, id)
	return err
}

const getBook = `-- name: GetBook :one
SELECT id, title, author_id, description, year_published, copies_total, copies_available
FROM books
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.AuthorID,
		&i.Description,
		&i.YearPublished,
		&i.CopiesTotal,
		&i.CopiesAvailable,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
SELECT
books.id,
books.author_id,
books.copies_available,
books.copies_total,
books.description,
books.title,
books.year_published,
authors.name
FROM books
JOIN authors on books.author_id = authors.id
ORDER BY title
`

type ListBooksRow struct {
	ID              int64
	AuthorID        int64
	CopiesAvailable int32
	CopiesTotal     int32
	Description     string
	Title           string
	YearPublished   int16
	Name            string
}

func (q *Queries) ListBooks(ctx context.Context) ([]ListBooksRow, error) {
	rows, err := q.db.QueryContext(ctx, listBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListBooksRow
	for rows.Next() {
		var i ListBooksRow
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.CopiesAvailable,
			&i.CopiesTotal,
			&i.Description,
			&i.Title,
			&i.YearPublished,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const partialUpdateBook = `-- name: PartialUpdateBook :one
UPDATE books
SET title = CASE WHEN $1::boolean THEN $2::VARCHAR(32) ELSE title END,
    description  = CASE WHEN $3::boolean THEN $4::TEXT ELSE description END
WHERE id = $5
RETURNING id, title, author_id, description, year_published, copies_total, copies_available
`

type PartialUpdateBookParams struct {
	UpdateTitle       bool
	Title             string
	UpdateDescription bool
	Description       string
	ID                int64
}

func (q *Queries) PartialUpdateBook(ctx context.Context, arg PartialUpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, partialUpdateBook,
		arg.UpdateTitle,
		arg.Title,
		arg.UpdateDescription,
		arg.Description,
		arg.ID,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.AuthorID,
		&i.Description,
		&i.YearPublished,
		&i.CopiesTotal,
		&i.CopiesAvailable,
	)
	return i, err
}

const truncateBooks = `-- name: TruncateBooks :exec
TRUNCATE books
`

func (q *Queries) TruncateBooks(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, truncateBooks)
	return err
}

const updateBook = `-- name: UpdateBook :one
UPDATE books
SET title = $2,
    description  = $3,
    author_id  = $4
WHERE id = $1
RETURNING id, title, author_id, description, year_published, copies_total, copies_available
`

type UpdateBookParams struct {
	ID          int64
	Title       string
	Description string
	AuthorID    int64
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBook,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.AuthorID,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.AuthorID,
		&i.Description,
		&i.YearPublished,
		&i.CopiesTotal,
		&i.CopiesAvailable,
	)
	return i, err
}
