package ask

import (
	"github.com/gin-gonic/gin"
)

func AskRouter(router *gin.Engine, askHandler *AskHandler) {
	router.POST("/ask", askHandler.AskHandler)
}
