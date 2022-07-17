package models

import (
	"database/sql"
)

type Book struct {
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

// Create a custom BookModel type which wraps the sql.DB connection pool.
type BookModel struct {
	DB *sql.DB
}

func (m BookModel) GetAll() ([]Book, error) {
	rows, err := m.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks []Book

	for rows.Next() {
		var bk Book

		if err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price); err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}

func (m BookModel) GetByIsbn(isbn string) (Book, error) {
	var bk Book

	row := m.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)
	if err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price); err != nil {
		return bk, err
	}
	return bk, nil
}

func (m BookModel) Update(isbn string, book Book) error {
	result, err := m.DB.Exec(`
		UPDATE books 
		SET isbn = $1, title = $2, author = $3, price = $4
		WHERE isbn = $5`, book.Isbn, book.Title, book.Author, book.Price, isbn,
	)

	if err != nil {
		return err
	}

	if _, err := result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (m BookModel) Post(book Book) (Book, error) {
	_, err := m.DB.Exec(`
		INSERT INTO books (isbn, title, artist, price)
		VALUES ($1, $2, $3, $4)`,
		book.Isbn, book.Title, book.Author, book.Price,
	)

	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func (m BookModel) Delete(book Book) error {
	_, err := m.DB.Exec("DELETE FROM books WHERE isbn = $1, title = $2, author = $3, price = $4", book.Isbn, book.Title, book.Author, book.Price)
	if err != nil {
		return err
	}
	return nil
}
