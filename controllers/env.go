package controllers

import "github.com/Jaim010/bookstore/models"

type Env struct {
	Books interface {
		GetAll() ([]models.Book, error)
		GetByIsbn(isbn string) (models.Book, error)
	}
}
