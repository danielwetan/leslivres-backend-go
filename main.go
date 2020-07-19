package main

import (
	"github.com/danielwetan/leslivres-backend-go/controllers"
	"github.com/danielwetan/leslivres-backend-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.Connect()

	r.GET("book/", controllers.GetBooks)
	r.GET("book/:id", controllers.GetBook)
	r.POST("book/", controllers.CreateBook)
	r.PUT("book/:id", controllers.UpdateBook)
	r.DELETE("book/:id", controllers.DeleteBook)
	r.Run()
}
