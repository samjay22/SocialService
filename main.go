package main

import (
	"github.com/heroku/go-getting-started/database"
	"github.com/heroku/go-getting-started/handler"
	"github.com/heroku/go-getting-started/services"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type Handlers struct {
	UserHandler handler.UserHandler
}

type Services struct {
	userService services.UserService
}

func InitServices() *Services {
	databaseObject := database.NewMockDatabase()

	userService := services.NewUserService(databaseObject)

	return &Services{
		userService: userService,
	}
}
func InitHandlers(services *Services, router *gin.Engine) *Handlers {

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
