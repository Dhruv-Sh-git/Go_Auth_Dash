package main

import (
	"log"

	"auth-dashboard-backend/config"
	"auth-dashboard-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize in-memory store
	config.InitStore()
	log.Println("In-memory user store initialized")

	// Initialize Gin router
	router := gin.Default()

	// Enable CORS for frontend
	router.Use(corsMiddleware())

	// Setup routes
	routes.SetupRoutes(router)

	// Welcome endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Auth Dashboard Backend API",
			"version": "1.0.0",
			"endpoints": gin.H{
				"health":   "GET /health",
				"register": "POST /api/auth/register",
				"login":    "POST /api/auth/login",
				"me":       "GET /api/user/me (requires JWT token)",
			},
		})
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Auth Dashboard Backend API",
		})
	})

	// Start server
	port := config.Env.Port
	log.Printf("Server starting on port %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// corsMiddleware enables CORS for frontend
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
