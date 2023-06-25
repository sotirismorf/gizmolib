package books

import (
	"context"
	"database/sql"
	"net/http"
	
	"github.com/sotirismorf/microservice/internal/database"
	"github.com/gin-gonic/gin"
)

type Service struct {
	queries *database.Queries
}

func NewService(queries *database.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.POST("/books", s.Create)
	router.GET("/books/:id", s.Get)
	router.PUT("/books/:id", s.FullUpdate)
	router.PATCH("/books/:id", s.PartialUpdate)
	router.DELETE("/books/:id", s.Delete)
	router.GET("/books", s.List)
}

type apiBook struct {
	ID          int64
	Title       string `json:"title,omitempty" binding:"required,max=32"`
	Description string `json:"description,omitempty" binding:"required"`
	AuthorID    int64
	YearPublished int16
	CopiesAvailable int32
	CopiesTotal int32
}

type apiBookFull struct {
	ID          int64
	AuthorID    int64
	Title       string `json:"title,omitempty" binding:"required,max=32"`
	AuthorName  string `json:"name,omitempty" binding:"required"`
	Description string `json:"description,omitempty" binding:"required"`
	YearPublished int16
	CopiesAvailable int32
	CopiesTotal int32
}

type apiBookPartialUpdate struct {
	Title       *string `json:"title,omitempty" binding:"omitempty,max=64"`
	Description *string `json:"description,omitempty" binding:"omitempty"`
}

func fromDB(book database.Book) *apiBook {
	return &apiBook{
		ID:          book.ID,
		AuthorID:    book.AuthorID,
		Title:       book.Title,
		Description: book.Description,
		YearPublished:    book.YearPublished,
		CopiesAvailable: book.CopiesAvailable,
		CopiesTotal: book.CopiesTotal,
	}
}

func fromDBFull(book database.ListBooksRow) *apiBookFull {
	return &apiBookFull{
		ID:          book.ID,
		AuthorID:    book.AuthorID,
		Title:       book.Title,
		Description: book.Description,
		AuthorName:  book.Name,
		YearPublished: book.YearPublished,
		CopiesTotal: book.CopiesTotal,
		CopiesAvailable: book.CopiesAvailable,
	}
}

type pathParameters struct {
	ID int64 `uri:"id" binding:"required"`
}

func (s *Service) Create(c *gin.Context) {
	// Parse request
	var request apiBook
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	params := database.CreateBookParams{
		Title:       request.Title,
		Description: request.Description,
		AuthorID:    request.AuthorID,
		YearPublished:    request.YearPublished,
		CopiesAvailable: request.CopiesAvailable,
		CopiesTotal: request.CopiesTotal,
	}
	book, err := s.queries.CreateBook(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	response := fromDB(book)
	c.IndentedJSON(http.StatusCreated, response)
}

func (s *Service) Get(c *gin.Context) {
	// Parse request
	var pathParams pathParameters
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get book
	book, err := s.queries.GetBook(context.Background(), pathParams.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	response := fromDB(book)
	c.IndentedJSON(http.StatusOK, response)
}

func (s *Service) FullUpdate(c *gin.Context) {
	// Parse request
	var pathParams pathParameters
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var request apiBook
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update book book
	params := database.UpdateBookParams{
		ID:   pathParams.ID,
		Title: request.Title,
		Description: request.Description,
		AuthorID: request.AuthorID,
	}
	book, err := s.queries.UpdateBook(context.Background(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	response := fromDB(book)
	c.IndentedJSON(http.StatusOK, response)
}

func (s *Service) PartialUpdate(c *gin.Context) {
	// Parse request
	var pathParams pathParameters
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var request apiBookPartialUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update book
	params := database.PartialUpdateBookParams{ID: pathParams.ID}
	if request.Title != nil {
		params.UpdateTitle = true
		params.Title = *request.Title
	}
	if request.Description != nil {
		params.UpdateDescription = true
		params.Description = *request.Description
	}
	book, err := s.queries.PartialUpdateBook(context.Background(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	response := fromDB(book)
	c.IndentedJSON(http.StatusOK, response)
}

func (s *Service) Delete(c *gin.Context) {
	// Parse request
	var pathParams pathParameters
	if err := c.ShouldBindUri(&pathParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete book
	if err := s.queries.DeleteBook(context.Background(), pathParams.ID); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	// Build response
	c.Status(http.StatusOK)
}

func (s *Service) List(c *gin.Context) {
	// List books
	books, err := s.queries.ListBooks(context.Background())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	if len(books) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Build response
	var response []*apiBookFull
	for _, book := range books {
		response = append(response, fromDBFull(book))
	}
	c.IndentedJSON(http.StatusOK, books)
}
