package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samjay22/SocialService/services"
	"github.com/samjay22/SocialService/structs"

	jsoniter "github.com/json-iterator/go"
)

type authHandler struct {
	authService services.AuthService
}

type AuthHandler interface {
	RegisterRoutes(router *gin.Engine)
}

func (s *authHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/login", s.login)
	router.POST("/logout", s.logout)
}

// VERIFY session id valid
func (s *authHandler) VerifySession(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		context.AbortWithError(404, err)
		return
	}

	requestData := structs.VerifySessionRequest{}
	err = jsoniter.Unmarshal(body, &requestData)
	if err != nil {
		context.AbortWithError(400, err)
		return
	}

	okay, err := s.authService.ValidateToken(context, requestData.SessionID)
	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"valid": okay})
}

// login
func (s *authHandler) login(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		context.AbortWithError(404, err)
		return
	}

	requestData := structs.LoginRequest{}
	err = jsoniter.Unmarshal(body, &requestData)
	if err != nil {
		context.AbortWithError(400, err)
		return
	}

	id, err := s.authService.Login(context, requestData.Username, requestData.Password)
	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "id": id})
}

// logout
func (s *authHandler) logout(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		context.AbortWithError(404, err)
		return
	}

	requestData := structs.LogoutRequest{}
	err = jsoniter.Unmarshal(body, &requestData)
	if err != nil {
		context.AbortWithError(400, err)
		return
	}

	err = s.authService.Logout(context, requestData.SessionID)
	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// constructor
func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}
