package routes

import (
	"document-management/core/config"
	"document-management/plugins/workspaces_plugin/controllers"
	"document-management/plugins/workspaces_plugin/repository"
	"document-management/plugins/workspaces_plugin/service"
	"document-management/share/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpWorkspaceRoutes(r *gin.Engine) {
	workspaceRepo := repository.NewWorkSpacesRepo(config.ConnectDB())
	workspaceService := service.NewWorkspaceService(workspaceRepo)
	workspaceController := controllers.NewWorkspaceController(workspaceService)

	workspace := r.Group("/api/v1/workspace", middleware.RequireAuth)
	{
		workspace.POST("/create", workspaceController.CreateWorkspace)
	}

}
