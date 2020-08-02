package controllers

import (
	"net/http"

	"github.com/danielwetan/leslivres-backend-go/src/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	register := models.Users{Username: input.Username, FullName: input.FullName, Email: input.Email, Password: input.Password, Role: input.Role}
	models.DB.Create(&register)

	c.JSON(http.StatusOK, gin.H{"data": register})
}

func Login(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	login := models.Users{Username: input.Username, Password: input.Password}
	models.DB.Where("username = ?", login).Find(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}
