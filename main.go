package main

import (
	"github.com/danielwetan/leslivres-backend-go/src/controllers"
	"github.com/danielwetan/leslivres-backend-go/src/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/danielwetan/leslivres-backend-go/src/routes"
)

func main() {
	r := gin.Default()
	// dotenv
	godotenv.Load()
	// name := os.Getenv("APP_NAME")
	// fmt.Println("hello", name)

	models.Connect()
	
	book := r.Group("/book")
	{
		book.GET("/", controllers.GetBooks)
		book.GET("/:id", controllers.GetBook)
		book.POST("/", controllers.CreateBook)
		book.PUT("/:id", controllers.UpdateBook)
		book.DELETE("/:id", controllers.DeleteBook)
	}

	author := r.Group("/author")
	{
		author.GET("/", controllers.GetAuthors)
		author.POST("/", controllers.CreateAuthor)
		author.PUT("/:id", controllers.UpdateAuthor)
		author.DELETE("/:id", controllers.DeleteAuthor)
	}

	r.Run()
}
