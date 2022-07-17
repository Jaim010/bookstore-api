package main

import (
	"log"

	"github.com/Jaim010/bookstore/controllers"
	"github.com/Jaim010/bookstore/database"
	"github.com/Jaim010/bookstore/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Init()
	if err != nil {
		log.Fatal("Database connection could not be made, exiting programn")
	}

	env := &controllers.Env{
		Books: models.BookModel{DB: db},
	}

	router := gin.Default()
	router.GET("/books/", env.GetAllBooks)
	router.GET("/books/:isbn", env.GetBook)

	router.Run("localhost:8080")
}
