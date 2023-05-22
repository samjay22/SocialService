package main

import (
	"os"

	"github.com/samjay22/SocialService/database"
	"github.com/samjay22/SocialService/handler"
	"github.com/samjay22/SocialService/services"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type Handlers struct {
	UserHandler handler.UserHandler
}

type Services struct {
	userService services.UserService
	authService services.AuthService
}

func InitServices() *Services {
	db, err := database.NewMockDatabase("redis://default:Vem71YeMHS3JqSpsYF17YhfXZHh4FKVg@redis-11523.c276.us-east-1-2.ec2.cloud.redislabs.com:11523")
	if err != nil {
		return nil
	}

	userService := services.NewUserService(db)
	authService := services.NewAuthService()

	return &Services{
		userService: userService,
		authService: authService,
	}
}

func InitHandlers(services *Services, router *gin.Engine) *Handlers {

	authService := handler.NewAuthHandler(services.authService)
	authService.RegisterRoutes(router)

	userHandler := handler.NewUserHandler(services.userService)
	userHandler.RegisterRoutes(router)

	return &Handlers{
		UserHandler: userHandler,
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	services := InitServices()
	InitHandlers(services, router)

	router.Run(":" + port)
}
