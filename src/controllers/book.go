package controllers

import (
	"net/http"

	"github.com/danielwetan/leslivres-backend-go/src/models"
	"github.com/gin-gonic/gin"
)

// Get all books
func GetBooks(c *gin.Context) {
	var books []models.Book
	// models.DB.Find(&books)
	// models.DB.Joins("INNER JOIN authors ON books.author = authors.id").Find(&books)
	// models.DB.Table("books").Select("books.id as id, books.title as title, books.description as description, books.img as img, authors.name as author").Joins("INNER JOIN authors ON books.author = authors.id").Scan(&books)
	models.DB.Table("books").Select("books.id as id, books.title as title, books.description as description, books.img as img, authors.name as author, genres.name as genre, books.status as status, books.date_added as added, books.date_updated as updated").Joins("INNER JOIN authors ON books.author = authors.id INNER JOIN genres ON books.genre = genres.id").Limit(12).Scan(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Get single book
func GetBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Create new book
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book := models.Book{Title: input.Title, Description: input.Description, Author: input.Author, Genre: input.Genre, Status: input.Status, Img: input.Img}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Update book
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete book
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": "Book successfully deleted"})
}
