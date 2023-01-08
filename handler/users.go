package handler

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/samjay22/SocialService/services"
)

type userHandler struct {
	userService services.UserService
}

type UserHandler interface {
}

func (s *userHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/users", s.GetUser)
}

type GetUserRequest struct {
	UserId string `json:"userId"`
}

func (s *userHandler) GetUser(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		context.AbortWithError(404, err)
		return
	}

	userData := GetUserRequest{}
	err = jsoniter.Unmarshal(body, &userData)
	if err != nil {
		context.AbortWithError(400, err)
		return
	}

	user, err := s.userService.GetUserByID(userData.UserId)
	if err != nil {
		context.AbortWithError(404, err)
		return
	}

	response, err := jsoniter.Marshal(user)
	if err != nil {
		context.AbortWithError(500, err)
	}

	context.Writer.Write(response)
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}

}
