package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/okkyPratama/bookstore-api/repository"
	"github.com/okkyPratama/bookstore-api/structs"
)

var bookRepo *repository.BookRepository

func SetBookRepository(repo *repository.BookRepository) {
	bookRepo = repo
}

func GetAllBooks(c *gin.Context) {
	books, err := bookRepo.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": books})
}

func CreateBook(c *gin.Context) {
	var book structs.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release year must be between 1980 and 2024"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	if err := bookRepo.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := bookRepo.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book structs.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = id

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release year must be between 1980 and 2024"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	if err := bookRepo.UpdateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := bookRepo.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func GetBooksByCategory(c *gin.Context) {
	categoryID, _ := strconv.Atoi(c.Param("categoryID"))
	books, err := bookRepo.GetBooksByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": books})
}
