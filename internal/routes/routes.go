package routes

import (
	"github.com/gin-gonic/gin"

	"go-user-api/internal/handlers"
	"go-user-api/internal/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)

	protected := api.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/users", handlers.GetUsers)
		protected.GET("/users/:id", handlers.GetUser)
		protected.PUT("/users/:id", handlers.UpdateUser)
		protected.DELETE("/users/:id", handlers.DeleteUser)
	}
}
