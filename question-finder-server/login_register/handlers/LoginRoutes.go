package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, loginHandler *LoginRegisterHandler) {
	router.POST("/register", loginHandler.RegisterHandler)
	router.POST("/login", loginHandler.LoginHandler)
	router.POST("/google-login", loginHandler.GoogleLoginHandler)
}
