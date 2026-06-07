package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/service"
)

type WorkoutHandler struct {
	workoutService *service.WorkoutService
}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{
		workoutService: service.NewWorkoutService(),
	}
}

func (h *WorkoutHandler) CreateTemplate(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req dto.WorkoutTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := h.workoutService.CreateTemplate(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *WorkoutHandler) GetAllTemplates(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	resp, err := h.workoutService.GetAllTemplates(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *WorkoutHandler) GetTemplate(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	templateID := c.Param("id")

	resp, err := h.workoutService.GetTemplate(userID, templateID)
	if err != nil {
		if err.Error() == "akses ditolak" {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *WorkoutHandler) UpdateTemplate(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	templateID := c.Param("id")

	var req dto.WorkoutTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := h.workoutService.UpdateTemplate(userID, templateID, req)
	if err != nil {
		if err.Error() == "akses ditolak" {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *WorkoutHandler) DeleteTemplate(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	templateID := c.Param("id")

	err := h.workoutService.DeleteTemplate(userID, templateID)
	if err != nil {
		if err.Error() == "akses ditolak" {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "template berhasil dihapus"})
}