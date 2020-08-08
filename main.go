package main

import (
	"github.com/danielwetan/leslivres-backend-go/src/controllers"
	"github.com/danielwetan/leslivres-backend-go/src/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	// "github.com/danielwetan/leslivres-backend-go/src/routes"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()  // router

	// Serve static files
	r.Static("/img", "./src/assets")

	// Cors
	r.Use(cors.Default())
	// dotenv
	godotenv.Load()
	PORT := os.Getenv("PORT")

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

	genre := r.Group("/genre")
	{
		genre.GET("/", controllers.GetGenres)
		genre.POST("/", controllers.CreateGenre)
		genre.PUT("/:id", controllers.UpdateGenre)
		genre.DELETE("/:id", controllers.DeleteGenre)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	user := r.Group("/user")
	{
		user.GET("/:id", controllers.GetUserData)
	}

	transaction := r.Group("/transaction")
	{
		transaction.GET("/:id", controllers.GetTranscation)
		transaction.POST("/", controllers.CreateTransaction)
		transaction.PUT("/:id", controllers.UpdateTransaction)
		transaction.DELETE("/:id", controllers.DeleteTransaction)
	}

	r.Run(PORT)
}
