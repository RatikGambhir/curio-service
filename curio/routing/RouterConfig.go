package routing

import (
	"curio/app"
	"curio/chat"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, config *app.AppConfig) {
	chat.RegisterChatRoutes(router, config)
}
