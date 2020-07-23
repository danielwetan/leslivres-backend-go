package main

import (
	"github.com/danielwetan/leslivres-backend-go/src/controllers"
	"github.com/danielwetan/leslivres-backend-go/src/models"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// dotenv
	godotenv.Load()
	// name := os.Getenv("APP_NAME")
	// fmt.Println("hello", name)

	models.Connect()

	r.GET("book/", controllers.GetBooks)
	r.GET("book/:id", controllers.GetBook)
	r.POST("book/", controllers.CreateBook)
	r.PUT("book/:id", controllers.UpdateBook)
	r.DELETE("book/:id", controllers.DeleteBook)

	r.GET("author/", controllers.GetAuthors)
	r.POST("author/", controllers.CreateAuthor)
	r.PUT("author/:id", controllers.UpdateAuthor)
	r.DELETE("author/:id", controllers.DeleteAuthor)
	r.Run()
}
