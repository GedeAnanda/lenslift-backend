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
		api.GET("/me", func(c *gin.Context) {
			userID := c.MustGet("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})

		profileHandler := NewProfileHandler()
		api.GET("/profile", profileHandler.GetProfile)
		api.PUT("/profile", profileHandler.UpdateProfile)

		workoutHandler := NewWorkoutHandler()
		api.POST("/workout-templates", workoutHandler.CreateTemplate)
		api.GET("/workout-templates", workoutHandler.GetAllTemplates)
		api.GET("/workout-templates/:id", workoutHandler.GetTemplate)
		api.PUT("/workout-templates/:id", workoutHandler.UpdateTemplate)
		api.DELETE("/workout-templates/:id", workoutHandler.DeleteTemplate)

		scheduleHandler := NewScheduleHandler()
		api.POST("/schedules", scheduleHandler.SetSchedule)
		api.GET("/schedules", scheduleHandler.GetSchedules)
		api.DELETE("/schedules/:day", scheduleHandler.DeleteSchedule)

		sessionHandler := NewSessionHandler()
		api.POST("/sessions/start", sessionHandler.StartSession)
		api.POST("/sessions/:id/log", sessionHandler.LogSet)
		api.POST("/sessions/:id/end", sessionHandler.EndSession)
		api.GET("/sessions", sessionHandler.GetAllSessions)
		api.GET("/sessions/:id", sessionHandler.GetSession)

		foodHandler := NewFoodHandler()
		api.POST("/food-logs", foodHandler.AddFoodLog)
		api.POST("/food-logs/analyze", foodHandler.AnalyzeFood)
		api.GET("/food-logs", foodHandler.GetDailyLogs)
		api.DELETE("/food-logs/:id", foodHandler.DeleteFoodLog)
	}

	return r
}