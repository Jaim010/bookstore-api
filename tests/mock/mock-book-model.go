package mock

import (
	"database/sql"
	"fmt"

	"github.com/Jaim010/bookstore/pkg/models"
)

type MockBookModel struct{}

func (m *MockBookModel) GetAll() ([]models.Book, error) {
	var bks = []models.Book{
		{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44},
		{Isbn: "978-1505255607", Title: "The Time Machine", Author: "H. G. Wells", Price: 5.99},
	}

	return bks, nil
}

func (m *MockBookModel) GetByIsbn(isbn string) (models.Book, error) {
	var bks = []models.Book{
		{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44},
		{Isbn: "978-1505255607", Title: "The Time Machine", Author: "H. G. Wells", Price: 5.99},
	}

	for _, bk := range bks {
		if bk.Isbn == isbn {
			return bk, nil
		}
	}

	return models.Book{}, sql.ErrNoRows
}

func (m *MockBookModel) Update(isbn string, book models.Book) error {
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

func (m *MockBookModel) Post(book models.Book) (models.Book, error) {
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

func (m *MockBookModel) Delete(book models.Book) error {
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
