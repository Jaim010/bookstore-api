package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Jaim010/bookstore/models"
	"github.com/gin-gonic/gin"
)

func (env *Env) GetAllBooks(c *gin.Context) {
	bks, err := env.Books.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, bks)
	}
}

func (env *Env) GetBook(c *gin.Context) {
	isbn := c.Param("isbn")
	bk, err := env.Books.GetByIsbn(isbn)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, bk)
}

func (env *Env) PutBook(c *gin.Context) {
	var newBook models.Book

	isbn := c.Param("isbn")
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isbn != newBook.Isbn {
		err := fmt.Sprintf("URI isbn: '%s' not equal to book isbn: '%s'", isbn, newBook.Isbn)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := env.Books.Update(isbn, newBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (env *Env) PostBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := env.Books.Post(newBook)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, book)
}

func (env *Env) DeleteBook(c *gin.Context) {
	isbn := c.Param("isbn")

	book, err := env.Books.GetByIsbn(isbn)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = env.Books.Delete(book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
