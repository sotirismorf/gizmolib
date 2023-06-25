package model

type ApiBookPartialUpdate struct {
	Title       *string `json:"title,omitempty" binding:"omitempty,max=64"`
	Description *string `json:"description,omitempty" binding:"omitempty"`
}

type ApiBook struct {
	ID              int64
	Title           string `json:"title,omitempty" binding:"required,max=32"`
	Description     string `json:"description,omitempty" binding:"required"`
	AuthorID        int64
	YearPublished   int16
	CopiesAvailable int32
	CopiesTotal     int32
}

type ApiBookFull struct {
	ID              int64
	AuthorID        int64
	Title           string `json:"title,omitempty" binding:"required,max=32"`
	AuthorName      string `json:"name,omitempty" binding:"required"`
	Description     string `json:"description,omitempty" binding:"required"`
	YearPublished   int16
	CopiesAvailable int32
	CopiesTotal     int32
}
