package routes

import (
	"document-management/core/config"
	"document-management/plugins/auth-plugin/controllers"
	"document-management/plugins/auth-plugin/repository"
	"document-management/plugins/auth-plugin/services"
	"github.com/gin-gonic/gin"
)

func SetAuthUpRoutes(r *gin.Engine) {

	userRepo := repository.NewUserRepository(config.ConnectDB())
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	auth := r.Group("api/v1/auth")
	{
		auth.POST("/login", userController.Login)
	}
}
