package controllers

import (
	"net/http"

	"github.com/danielwetan/leslivres-backend-go/src/models"
	"github.com/gin-gonic/gin"
)

// Get Transaction
func GetTranscation(c *gin.Context) {
	var transaction models.Transaction

	if err := models.DB.Where("id=?", c.Param("id")).First(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

// Create new transaction
func CreateTransaction(c *gin.Context) { 
	var input models.CreateTransactioniInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	transaction := models.Transaction{User: input.User, Book: input.Book, Status: input.Status}
	models.DB.Create(&transaction)

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

// Update transaction
func UpdateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := models.DB.Where("id=?", c.Param("id")).First(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	// Validate input
	var input models.UpdateTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	models.DB.Model(&transaction).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

// Delete transaction
func DeleteTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := models.DB.Where("id=?", c.Param("id")).First(&transaction).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	models.DB.Delete(&transaction)

	c.JSON(http.StatusOK, gin.H{"data": "Transaction successfully deleted!"})
}