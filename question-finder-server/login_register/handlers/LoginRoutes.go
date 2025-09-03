package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, loginHandler *LoginRegisterHandler) {
	router.POST("/register", loginHandler.RegisterHandler)
	router.POST("/login", loginHandler.LoginHandler)
	router.GET("/google/login", loginHandler.GoogleLoginHandler)
	router.GET("/google/callback", loginHandler.GoogleLoginCallbackHandler)
}
