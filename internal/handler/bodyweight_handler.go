package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/service"
)

type BodyWeightHandler struct {
	bodyWeightService *service.BodyWeightService
}

func NewBodyWeightHandler() *BodyWeightHandler {
	return &BodyWeightHandler{
		bodyWeightService: service.NewBodyWeightService(),
	}
}

func (h *BodyWeightHandler) LogWeight(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req dto.BodyWeightRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := h.bodyWeightService.LogWeight(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *BodyWeightHandler) GetHistory(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	resp, err := h.bodyWeightService.GetHistory(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *BodyWeightHandler) GetLatest(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	resp, err := h.bodyWeightService.GetLatest(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *BodyWeightHandler) DeleteWeight(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	weightID := c.Param("id")

	err := h.bodyWeightService.DeleteWeight(userID, weightID)
	if err != nil {
		if err.Error() == "akses ditolak" {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berat badan berhasil dihapus"})
}