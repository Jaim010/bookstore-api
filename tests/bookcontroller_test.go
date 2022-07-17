package controllers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Jaim010/bookstore/pkg/controllers"
	"github.com/Jaim010/bookstore/pkg/models"
	"github.com/gin-gonic/gin"
)

type mockBookModel struct{}

func (m *mockBookModel) GetAll() ([]models.Book, error) {
	var bks = []models.Book{
		{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44},
		{Isbn: "978-1505255607", Title: "The Time Machine", Author: "H. G. Wells", Price: 5.99},
	}

	return bks, nil
}

func (m *mockBookModel) GetByIsbn(isbn string) (models.Book, error) {
	var bks = []models.Book{
		{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44},
		{Isbn: "978-1505255607", Title: "The Time Machine", Author: "H. G. Wells", Price: 5.99},
	}

	for _, bk := range bks {
		if bk.Isbn == isbn {
			return bk, nil
		}
	}

	return models.Book{}, fmt.Errorf("book not found")
}

func (m *mockBookModel) Update(isbn string, book models.Book) error {
	var bks = []models.Book{
		{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44},
		{Isbn: "978-1505255607", Title: "The Time Machine", Author: "H. G. Wells", Price: 5.99},
	}

	for i, bk := range bks {
		if bk.Isbn == isbn {
			bks[i] = book
			return nil
		}
	}

	return fmt.Errorf("book not found")
}

func (m *mockBookModel) Post(book models.Book) (models.Book, error) {
	var bks = []models.Book{
		{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44},
		{Isbn: "978-1505255607", Title: "The Time Machine", Author: "H. G. Wells", Price: 5.99},
	}

	for _, bk := range bks {
		if bk.Isbn == book.Isbn {
			return models.Book{}, fmt.Errorf("book with isbn %s already exists", book.Isbn)
		}
	}

	bks = append(bks, book)
	return book, nil
}

func (m *mockBookModel) Delete(book models.Book) error {
	var bks = []models.Book{
		{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44},
		{Isbn: "978-1505255607", Title: "The Time Machine", Author: "H. G. Wells", Price: 5.99},
	}

	for i, bk := range bks {
		if bk.Isbn == book.Isbn {
			bks = append(bks[:i], bks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("book not found")
}

func TestGetAllBooks(t *testing.T) {
	// Arrange
	expectedBooks := []models.Book{
		{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44},
		{Isbn: "978-1505255607", Title: "The Time Machine", Author: "H. G. Wells", Price: 5.99},
	}

	env := controllers.Env{Books: &mockBookModel{}}
	router := gin.Default()
	router.GET("/books/", env.GetAllBooks)
	req, _ := http.NewRequest("GET", "/books/", nil)
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	responseData, _ := ioutil.ReadAll(w.Body)
	var responseBooks []models.Book

	err := json.Unmarshal(responseData, &responseBooks)
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(responseBooks, expectedBooks) {
		t.Errorf("Output %v not equal to expected %v", responseBooks, expectedBooks)
	}
}
