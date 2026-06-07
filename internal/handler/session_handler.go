package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanda/lenslift-backend/internal/dto"
	"github.com/nanda/lenslift-backend/internal/service"
)

type SessionHandler struct {
	sessionService *service.SessionService
}

func NewSessionHandler() *SessionHandler {
	return &SessionHandler{
		sessionService: service.NewSessionService(),
	}
}

func (h *SessionHandler) StartSession(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req dto.StartSessionRequest
	c.ShouldBindJSON(&req)

	resp, err := h.sessionService.StartSession(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *SessionHandler) LogSet(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	sessionID := c.Param("id")

	var req dto.LogSetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := h.sessionService.LogSet(userID, sessionID, req)
	if err != nil {
		if err.Error() == "akses ditolak" {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *SessionHandler) EndSession(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	sessionID := c.Param("id")

	resp, err := h.sessionService.EndSession(userID, sessionID)
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

func (h *SessionHandler) GetAllSessions(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	resp, err := h.sessionService.GetAllSessions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if resp == nil {
		c.JSON(http.StatusOK, []dto.SessionResponse{})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *SessionHandler) GetSession(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	sessionID := c.Param("id")

	resp, err := h.sessionService.GetSession(userID, sessionID)
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