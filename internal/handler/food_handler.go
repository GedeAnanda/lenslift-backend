package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/service"
)

type FoodHandler struct {
	foodService *service.FoodService
}

func NewFoodHandler() *FoodHandler {
	return &FoodHandler{
		foodService: service.NewFoodService(),
	}
}

func (h *FoodHandler) AddFoodLog(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req dto.FoodLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := h.foodService.AddFoodLog(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *FoodHandler) AnalyzeFood(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "file gambar tidak ditemukan"})
		return
	}
	defer file.Close()

	resp, err := h.foodService.AnalyzeFood(userID, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *FoodHandler) GetDailyLogs(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	date := c.Query("date")

	resp, err := h.foodService.GetDailyLogs(userID, date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *FoodHandler) DeleteFoodLog(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	foodID := c.Param("id")

	err := h.foodService.DeleteFoodLog(userID, foodID)
	if err != nil {
		if err.Error() == "akses ditolak" {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "log makanan berhasil dihapus"})
}