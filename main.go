package main

import (
	"go-supabase-crud/database"
	"go-supabase-crud/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.DB.Close()

	// 2. Create Gin router
	router := gin.Default()

	// 3. Setup routes
	routes.SetupRoutes(router, database.DB)

	// 4. Start server
	log.Println("Server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
