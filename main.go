package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/okkyPratama/bookstore-api/controllers"
	"github.com/okkyPratama/bookstore-api/db"
	"github.com/okkyPratama/bookstore-api/middleware"
	"github.com/okkyPratama/bookstore-api/repository"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)
	os.Getenv("PORT")
	err = db.InitDB(psqlInfo)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.DbConnection.Close()

	err = db.DBMigrate()
	if err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	bookRepo := &repository.BookRepository{DB: db.DbConnection}
	categoryRepo := &repository.CategoryRepository{DB: db.DbConnection}
	userRepo := &repository.UserRepository{DB: db.DbConnection}

	controllers.SetBookRepository(bookRepo)
	controllers.SetCategoryRepository(categoryRepo)
	controllers.SetUserRepository(userRepo)

	r := gin.Default()

	// Public routes
	r.POST("/api/users/register", controllers.RegisterUser)
	r.POST("/api/users/login", controllers.LoginUser)

	// Protected routes
	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/books", controllers.GetAllBooks)
		authorized.POST("/books", controllers.CreateBook)
		authorized.GET("/books/:id", controllers.GetBookByID)
		authorized.PUT("/books/:id", controllers.UpdateBook)
		authorized.DELETE("/books/:id", controllers.DeleteBook)
		authorized.GET("/categories/:id/books", controllers.GetBooksByCategory)

		authorized.GET("/categories", controllers.GetAllCategories)
		authorized.POST("/categories", controllers.CreateCategory)
		authorized.GET("/categories/:id", controllers.GetCategoryByID)
		authorized.PUT("/categories/:id", controllers.UpdateCategory)
		authorized.DELETE("/categories/:id", controllers.DeleteCategory)
	}

	fmt.Println("Server is running on port 8080...")
	r.Run(os.Getenv("PORT"))
}
