package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda/lenslift-backend/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/auth")
	{
		authHandler := NewAuthHandler()
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		profileHandler := NewProfileHandler()
		api.GET("/profile", profileHandler.GetProfile)
		api.PUT("/profile", profileHandler.UpdateProfile)

		api.GET("/me", func(c *gin.Context) {
			userID := c.MustGet("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})
	}

	return r
}


