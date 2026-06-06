package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/service"
)

type ProfileHandler struct {
	profileService *service.ProfileService
}

func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{
		profileService: service.NewProfileService(),
	}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	resp, err :=h.profileService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "profil tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := h.profileService.UpdateProfile(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK,resp)
}
