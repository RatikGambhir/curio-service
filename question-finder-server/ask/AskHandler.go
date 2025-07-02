package ask

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AskHandler struct {
}

func ConstructAskHandler() *AskHandler {
	return &AskHandler{}
}

func (h *AskHandler) AskHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
