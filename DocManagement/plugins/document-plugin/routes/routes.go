package routes

import (
	"document-management/core/config"
	"document-management/plugins/document-plugin/controllers"
	"document-management/plugins/document-plugin/repository"
	"document-management/plugins/document-plugin/services"
	"document-management/share/middleware"
	"github.com/gin-gonic/gin"
)

func SetDocumentUpRoutes(r *gin.Engine) {

	docRepo := repository.NewDocumentRepository(config.ConnectDB())
	docService := services.NewDocumentService(docRepo)
	docController := controllers.NewDocumentController(docService)

	doc := r.Group("api/v1/doc", middleware.RequireAuth)
	{
		doc.POST("/upload", docController.UploadDocument)
		doc.GET("/get/:id", docController.GetDocumentByUserId)
		doc.GET("/download", docController.DownLoadController)
	}
}
