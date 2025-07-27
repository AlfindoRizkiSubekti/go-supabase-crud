package routes

import (
	"go-supabase-crud/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(router *gin.Engine, db *pgxpool.Pool) {
	// Initialize controllers
	bookController := controllers.NewBookController(db)
	authorController := controllers.NewAutorController(db)
	// ... tambahkan controller lainnya

	// Group routes
	api := router.Group("/api")
	{
		// Author routes
		api.GET("/authors", authorController.GetAuthors)

		// Book routes
		api.GET("/books", bookController.GetBooks)

		// ... tambahkan routes lainnya
	}
}
