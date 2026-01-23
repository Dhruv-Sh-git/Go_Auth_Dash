//api endpoints, business logic of apis are in controller
package routes

import (
	"auth-dashboard-backend/controllers"
	"auth-dashboard-backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine) {
	// API group
	api := router.Group("/api")
	{
		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// User routes (protected)
		user := api.Group("/user")
		user.Use(middleware.JWTAuth())
		{
			user.GET("/me", controllers.GetMe)
		}
	}
}
