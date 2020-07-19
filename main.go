package main

import (
	"github.com/danielwetan/leslivres-backend-go/controllers"
	"github.com/danielwetan/leslivres-backend-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.Connect()

	r.GET("books/", controllers.GetBooks)
	r.Run()
}
