package controllers

import "github.com/Jaim010/go-gin-bookstore-api/pkg/models"

type Env struct {
	Books interface {
		GetAll() ([]models.Book, error)
		GetByIsbn(isbn string) (models.Book, error)
		Update(isbn string, book models.Book) error
		Post(book models.Book) (models.Book, error)
		Delete(book models.Book) error
	}
}
