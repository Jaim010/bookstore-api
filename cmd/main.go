package main

import (
	"log"

	"github.com/Jaim010/go-gin-bookstore-api/pkg/config"
	"github.com/Jaim010/go-gin-bookstore-api/pkg/controllers"
	"github.com/Jaim010/go-gin-bookstore-api/pkg/database"
	"github.com/Jaim010/go-gin-bookstore-api/pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("Configuration failed, exiting programn")
	}
	db, err := database.Init()
	if err != nil {
		log.Fatal("Failed to connect to database, exiting programn")
	}

	env := &controllers.Env{
		Books: models.BookModel{DB: db},
	}

	router := gin.Default()

	router.GET("/health", controllers.GetHealth)

	router.GET("/books/", env.GetAllBooks)
	router.GET("/books/:isbn", env.GetBook)
	router.PUT("/books/:isbn", env.PutBook)
	router.POST("/books/", env.PostBook)
	router.DELETE("/books/:isbn", env.DeleteBook)

	router.Run("localhost:8080")
}
