package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sotirismorf/microservice/api/authors"
	"github.com/sotirismorf/microservice/api/books"
	"github.com/sotirismorf/microservice/api/token"
	"github.com/sotirismorf/microservice/cmd/microservice/config"
	"github.com/sotirismorf/microservice/internal/database"
	"golang.org/x/crypto/bcrypt"
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

	queries := database.New(postgres.DB)

	// Initialize admin user
	if _, err := queries.GetUserById(context.Background(), 1); err != nil {
		log.Println("No admin user found in database. Creating admin user...")

		if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cfg.Admin.Password), bcrypt.DefaultCost); err != nil {
			log.Println(err.Error())
		} else {
			params := database.CreateUserParams{
				Username: cfg.Admin.Username,
				Password: string(hashedPassword),
			}

			if _, err := queries.CreateUser(context.Background(), params); err != nil {
				log.Println("Something went wrong")
			} else {
				log.Println("Admin user created")
			}
		}
	} else {
		log.Println("Found admin user in database")
	}

	// Instantiates the author service
	authorService := authors.NewService(queries)
	bookService := books.NewService(queries)
	tokenService := token.NewService(queries)

	// Register our service handlers to the router
	router := gin.Default()
	authorService.RegisterHandlers(router)
	bookService.RegisterHandlers(router)
	tokenService.RegisterHandlers(router)

	// Start the server
	router.Run()
}
