package model

type ApiBookPartialUpdate struct {
	Title       *string `json:"title,omitempty" binding:"omitempty,max=64"`
	Description *string `json:"description,omitempty" binding:"omitempty"`
}

type ApiBook struct {
	ID              int64  `json:"id"`
	Title           string `json:"title,omitempty" binding:"required,max=32"`
	Description     string `json:"description,omitempty" binding:"required"`
	AuthorID        int64  `json:"authorId"`
	YearPublished   int16  `json:"yearPublished"`
	CopiesAvailable int32  `json:"copiesAvailable"`
	CopiesTotal     int32  `json:"copiesTotal"`
}

type ApiBookFull struct {
	ID              int64  `json:"id"`
	AuthorID        int64  `json:"authorId"`
	Title           string `json:"title,omitempty" binding:"required,max=32"`
	AuthorName      string `json:"authorName,omitempty" binding:"required"`
	Description     string `json:"description,omitempty" binding:"required"`
	YearPublished   int16  `json:"yearPublished"`
	CopiesAvailable int32  `json:"copiesAvailable"`
	CopiesTotal     int32  `json:"copiesTotal"`
}
