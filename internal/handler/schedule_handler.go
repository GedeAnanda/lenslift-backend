package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/service"
)

type ScheduleHandler struct {
	scheduleService *service.ScheduleService
}

func NewScheduleHandler() *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: service.NewScheduleService(),
	}
}

func (h *ScheduleHandler) SetSchedule(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req dto.ScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := h.scheduleService.SetSchedule(userID, req)
	if err != nil {
		if err.Error() == "akses ditolak" {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ScheduleHandler) GetSchedules(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	resp, err := h.scheduleService.GetSchedules(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if resp == nil {
		c.JSON(http.StatusOK, []dto.ScheduleResponse{})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	day := c.Param("day")

	err := h.scheduleService.DeleteSchedule(userID, day)
	if err != nil {
		if err.Error() == "akses ditolak" {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "jadwal berhasil dihapus"})
}