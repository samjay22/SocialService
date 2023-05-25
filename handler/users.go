package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samjay22/SocialService/services"
	"github.com/samjay22/SocialService/structs"

	jsoniter "github.com/json-iterator/go"
)

type userHandler struct {
	userService services.UserService
}

type UserHandler interface {
	RegisterRoutes(router *gin.Engine)
}

func (s *userHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/getUser", s.GetUser)
	router.POST("/createUser", s.CreateUser)
	router.POST("/addFriend", s.AddFriend)
}

// function to add a friend to a user
func (s *userHandler) AddFriend(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		context.AbortWithError(404, err)
		return
	}

	userData := structs.AddFriendRequest{}
	err = jsoniter.Unmarshal(body, &userData)
	if err != nil {
		context.AbortWithError(400, err)
		return
	}

	err = s.userService.AddUserAsFriend(context, userData.UserId, userData.FriendId)
	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Friend added successfully"})
}

func (s *userHandler) GetUser(context *gin.Context) {
	userIdStr := context.Query("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		context.AbortWithError(400, err) // Return a 400 Bad Request if the userId is not a valid int64
		return
	}

	userData := structs.GetUserRequest{
		UserId: userId,
	}

	user, err := s.userService.GetUserByID(context, userData.UserId)
	if err != nil {
		context.AbortWithError(404, err)
		return
	}

	response, err := jsoniter.Marshal(user)
	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	context.Writer.Write(response)
}

func (s *userHandler) CreateUser(context *gin.Context) {
	var userData structs.CreateUserProfileRequest

	body, err := context.GetRawData()
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(body, &userData)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Call your user service to create the user
	err = s.userService.CreateUser(context, &userData)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}

}
