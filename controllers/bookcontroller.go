package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (env *Env) GetAllBooks(c *gin.Context) {
	bks, err := env.Books.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
	} else {
		c.IndentedJSON(http.StatusOK, bks)
	}
}

func (env *Env) GetBook(c *gin.Context) {
	isbn := c.Param("isbn")
	bk, err := env.Books.GetByIsbn(isbn)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusOK, nil)
		} else {
			c.IndentedJSON(http.StatusInternalServerError, nil)
		}
	} else {
		c.IndentedJSON(http.StatusOK, bk)
	}
}
