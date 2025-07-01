package handlers

import (
	"question_finder/middleware"

	"github.com/gin-gonic/gin"
)

func ProfileSettingsRoutes(router *gin.Engine, profileSettingsHandler *ProfileSettingsHandler) {
	routerGroup := router.Group("/profile-settings", middleware.AuthMiddleware())
	routerGroup.GET("/", profileSettingsHandler.UpdateProfileSettings)
}
