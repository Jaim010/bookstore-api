package controllers_unit_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Jaim010/go-gin-bookstore-api/pkg/controllers"
	"github.com/Jaim010/go-gin-bookstore-api/pkg/models"
	"github.com/Jaim010/go-gin-bookstore-api/tests/mock"
	"github.com/gin-gonic/gin"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkEqual[K any](t *testing.T, expected, actual K) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v. Got %v\n", expected, actual)
	}
}

func TestGetAllBooks(t *testing.T) {
	// Arrange
	expectedBooks := []models.Book{
		{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44},
		{Isbn: "978-1505255607", Title: "The Time Machine", Author: "H. G. Wells", Price: 5.99},
	}

	env := controllers.Env{Books: &mock.MockBookModel{}}
	router := gin.Default()
	router.GET("/books/", env.GetAllBooks)
	req, _ := http.NewRequest("GET", "/books/", nil)

	// Act
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	checkResponseCode(t, http.StatusOK, w.Code)
	var responseBooks []models.Book

	json.Unmarshal(w.Body.Bytes(), &responseBooks)

	checkEqual(t, expectedBooks, responseBooks)
}

type getBookTest struct {
	givenIsbn    string
	expectedBook models.Book
	expectedCode int
}

var getBookTests = []getBookTest{
	{"978-1503261969", models.Book{Isbn: "978-1503261969", Title: "Emma", Author: "Jayne Austen", Price: 9.44}, 200},
	{"123-1123937012", models.Book{}, 404},
	{"121232133-1122", models.Book{}, 404},
	{"88098091238809", models.Book{}, 404},
	{"abcdefghijklmn", models.Book{}, 404},
	{"", models.Book{}, 404},
}

func TestGetBook(t *testing.T) {
	// Arrange
	env := controllers.Env{Books: &mock.MockBookModel{}}
	router := gin.Default()
	router.GET("/books/:isbn", env.GetBook)

	for _, test := range getBookTests {
		req, _ := http.NewRequest("GET", "/books/"+test.givenIsbn, nil)
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		checkResponseCode(t, test.expectedCode, w.Code)
		if test.expectedCode == 200 {
			var responseBook models.Book
			json.Unmarshal(w.Body.Bytes(), &responseBook)

			checkEqual(t, test.expectedBook, responseBook)
		}
	}
}

type putBookTest struct {
	isbn         string
	bookData     string
	expectedCode int
}

var putBookTests = []putBookTest{
	{"978-1503261969", `{"isbn": "978-1503261969", "title": "Moby Dick", "author": "Jack Black", "price": 13.22}`, 204},
	{"978-1503261969", `{"isbn": "978-1503123456", "title": "Moby Dick", "author": "Jack Black", "price": 13.22}`, 400},
	{"978-1503123456", `{"isbn": "978-1503261969", "title": "Moby Dick", "author": "Jack Black", "price": 13.22}`, 400},
}

func TestPutBook(t *testing.T) {
	// Arrange
	env := controllers.Env{Books: &mock.MockBookModel{}}
	router := gin.Default()
	router.PUT("/books/:isbn", env.PutBook)

	for _, test := range putBookTests {
		jsonStr := []byte(test.bookData)
		req, _ := http.NewRequest("PUT", "/books/"+test.isbn, bytes.NewBuffer(jsonStr))
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		checkResponseCode(t, test.expectedCode, w.Code)
	}
}

type postBookTest struct {
	bookData     string
	expectedCode int
}

var postBookTests = []postBookTest{
	{`{"isbn": "349-4902119203", "title": "Moby Dick", "author": "Jack Black", "price": 13.22}`, 201},
	// {`{"isbn": "349-4902119203", "title": "Moby Dick", "author": "Jack Black", "price": 13.22}`, 400},
}

func TestPostBook(t *testing.T) {
	// Arrange
	env := controllers.Env{Books: &mock.MockBookModel{}}
	router := gin.Default()
	router.POST("/books/", env.PostBook)

	for _, test := range postBookTests {
		jsonStr := []byte(test.bookData)
		req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(jsonStr))
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		checkResponseCode(t, test.expectedCode, w.Code)
	}
}

type deleteBookTest struct {
	Isbn         string
	expectedCode int
}

var deleteBookTests = []deleteBookTest{
	{"978-1503261969", 204},
	{"123-4567890123", 404},
}

func TestDeleteBook(t *testing.T) {
	// Arrange
	env := controllers.Env{Books: &mock.MockBookModel{}}
	router := gin.Default()
	router.DELETE("/books/:isbn", env.DeleteBook)

	for _, test := range deleteBookTests {
		req, _ := http.NewRequest("DELETE", "/books/"+test.Isbn, nil)
		w := httptest.NewRecorder()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		checkResponseCode(t, test.expectedCode, w.Code)
	}
}
