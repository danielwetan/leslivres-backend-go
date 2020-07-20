package controllers

import (
	"net/http"

	"github.com/danielwetan/leslivres-backend-go/models"
	"github.com/gin-gonic/gin"
)

// Get all authors
func GetAuthors(c *gin.Context) {
	var authors []models.Author
	models.DB.Find(&authors)

	c.JSON(http.StatusOK, gin.H{"data": authors})
}

func CreateAuthor(c *gin.Context) {
	var input models.CreateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	author := models.Author{Name: input.Name}
	models.DB.Create(&author)

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func UpdateAuthor(c *gin.Context) {
	var author models.Author
	if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	var input models.UpdateAuthorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	models.DB.Model(&author).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": author})
}

func DeleteAuthor(c *gin.Context) {
	var author models.Author
	if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	models.DB.Delete(&author)
	c.JSON(http.StatusOK, gin.H{"data": "Author successfully deleted"})
}
