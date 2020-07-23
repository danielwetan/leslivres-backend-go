package controllers

import (
	"net/http"

	"github.com/danielwetan/leslivres-backend-go/src/models"
	"github.com/gin-gonic/gin"
)

// Get all genres
func GetGenres(c *gin.Context) {
	var genres []models.Genre
	models.DB.Find(&genres)

	c.JSON(http.StatusOK, gin.H{"data": genres})
}

func CreateGenre(c *gin.Context) {
	var input models.GenreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	genre := models.Genre{Name: input.Name}
	models.DB.Create(&genre)

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

func UpdateGenre(c *gin.Context) {
	var genre models.Genre
	if err := models.DB.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
	}

	var input models.GenreInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	models.DB.Model(&genre).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": genre})
}

func DeleteGenre(c *gin.Context) {
	var genre models.Genre
	if err := models.DB.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	models.DB.Delete(&genre)
	c.JSON(http.StatusOK, gin.H{"data": "Genre successfully deleted"})
}
