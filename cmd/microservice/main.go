package main

import (
	"log"

	"github.com/sotirismorf/microservice/api/authors"
	"github.com/sotirismorf/microservice/api/books"
	"github.com/sotirismorf/microservice/cmd/microservice/config"
	"github.com/sotirismorf/microservice/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Read configuration
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Instantiates the database
	postgres, err := database.NewPostgres(cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Instantiates the author service
	queries := database.New(postgres.DB)
	authorService := authors.NewService(queries)
	bookService := books.NewService(queries)

	// Register our service handlers to the router
	router := gin.Default()
	authorService.RegisterHandlers(router)
	bookService.RegisterHandlers(router)

	// Start the server
	router.Run()
}
