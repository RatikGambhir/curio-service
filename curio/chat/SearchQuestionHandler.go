package chat

import (
	"curio/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchQuestionHandler struct {
	searchQuestionProcessor *SearchQuestionProcessor
}

func NewSearchQuestionHandler(config *app.AppConfig) *SearchQuestionHandler {
	return &SearchQuestionHandler{
		searchQuestionProcessor: NewSearchQuestionProcessor(config),
	}
}

func (h *SearchQuestionHandler) KeywordSearchHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
func (askHandler *SearchQuestionHandler) ProcessQuestionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
