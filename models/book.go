package models

import (
	"database/sql"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
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
