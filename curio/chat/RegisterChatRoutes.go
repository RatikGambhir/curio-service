package chat

import (
	"curio/app"

	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(router *gin.Engine, config *app.AppConfig) {
	var searchQuestionHadler = NewSearchQuestionHandler(config)
	router.GET("/keyword-search", searchQuestionHadler.KeywordSearchHandler)
	router.POST("/process-question", searchQuestionHadler.ProcessQuestionHandler)
}
