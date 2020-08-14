package controllers

import (
	"net/http"

	"unsafe"

	"github.com/danielwetan/leslivres-backend-go/src/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	register := models.Users{Username: input.Username, FullName: input.FullName, Email: input.Email, Password: input.Password, Role: input.Role}
	// register.Password = "New password" // placeholder for hashing password
	// fmt.Println(register.Password)
	hash, _ := HashPassword(register.Password)
	register.Password = hash

	// fmt.Println("Check password 1:", CheckPasswordHash(register.Password, hash))
	// fmt.Println("Check password 2:", CheckPasswordHash("999", hash))
	// fmt.Println(hash);

	models.DB.Create(&register)

	c.JSON(http.StatusOK, gin.H{"data": register})
}

func Login(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	login := models.Users{Username: input.Username}
	result := models.DB.Where("username = ?", login.Username).First(&input)

	if unsafe.Sizeof(result) > 5 {
		c.JSON(http.StatusOK, gin.H{"data": result.Value})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "404 data not found"})
	}

}
