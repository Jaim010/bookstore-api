package controllers_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Jaim010/bookstore/pkg/controllers"
	"github.com/Jaim010/bookstore/pkg/models"
	"github.com/Jaim010/bookstore/tests/mock"
	"github.com/gin-gonic/gin"
)

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
