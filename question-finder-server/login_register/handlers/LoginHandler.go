package handlers

import (
	"log"
	"net/http"
	AppConfig "question_finder/app"
	"question_finder/login_register/processor"
	"question_finder/login_register/types"

	"github.com/gin-gonic/gin"
)

type LoginRegisterHandler struct {
	LoginProcessor *processor.LoginProcessor
}

func ConstructLoginRegisterHandler(appConfig *AppConfig.PostgresConfig) *LoginRegisterHandler {
	return &LoginRegisterHandler{
		LoginProcessor: processor.ConstructLoginProcessor(appConfig),
	}
}

func (h *LoginRegisterHandler) LoginHandler(c *gin.Context) {
	var req types.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})

	}
	resp, err := h.LoginProcessor.LoginUser(req)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": resp})
}

func (loginHandler *LoginRegisterHandler) RegisterHandler(c *gin.Context) {
	var req types.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Register Handler failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// TODO: Add email verification and account input logic here

	resp, err := loginHandler.LoginProcessor.RegisterUser(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user, please try again"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp})

}

func (loginHandler *LoginRegisterHandler) GoogleLoginHandler(c *gin.Context) {
	var googleLoginRequest = types.GoogleLoginRequest{}
	if err := c.ShouldBindJSON(&googleLoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

}
