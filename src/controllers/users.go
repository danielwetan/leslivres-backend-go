package controllers

import (
	// "log"
	"net/http"

	"github.com/danielwetan/leslivres-backend-go/src/models"
	"github.com/gin-gonic/gin"
)

// Get User Data
func GetUserData(c *gin.Context) {
	var data models.Users

	if err := models.DB.Where("id=?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	// log.Print("data :", data.Username)
	// username := "username" + data.Username
	// fullName := "full_name"+ ":" + data.FullName
	// email := "email:" + data.Email

	// msg := []string {username, fullName , email}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
