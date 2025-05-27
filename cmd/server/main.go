// cmd/server/main.go
package main

import (
	"log"
	"os"

	_ "github.com/alishercodecrafter/orderpackscalculator/docs" // Import generated docs
	"github.com/alishercodecrafter/orderpackscalculator/internal/controller"
	"github.com/alishercodecrafter/orderpackscalculator/internal/repository"
	"github.com/alishercodecrafter/orderpackscalculator/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Create repository, service, and controller
	repo := repository.NewMemoryRepository()
	svc := service.NewPacksService(repo)
	ctrl := controller.NewPacksController(svc)

	// Create Gin router
	router := gin.Default()

	// Load HTML templates and static files
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "web/static")

	// Define routes
	router.GET("/", ctrl.GetIndex)

	// API routes
	api := router.Group("/api")
	{
		api.GET("/packs", ctrl.GetPacks)
		api.POST("/packs", ctrl.AddPack)
		api.DELETE("/packs/:size", ctrl.RemovePack)
		api.POST("/calculate", ctrl.CalculatePacks)
	}

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	log.Printf("Swagger documentation available at http://localhost:%s/swagger/index.html", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
