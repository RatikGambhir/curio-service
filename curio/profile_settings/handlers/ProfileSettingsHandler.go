package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileSettingsHandler struct {
}

func ConstructProfileSettingsHandler() *ProfileSettingsHandler {
	return &ProfileSettingsHandler{}
}

func (h *ProfileSettingsHandler) UpdateProfileSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Profile settings updated"})
}
