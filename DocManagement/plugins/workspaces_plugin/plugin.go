package workspaces_plugin

import (
	"document-management/plugins/workspaces_plugin/routes"
	"github.com/gin-gonic/gin"
)

type WorkspacePlugin struct{}

func NewDocumentPlugin() *WorkspacePlugin {
	return &WorkspacePlugin{}
}

func (w *WorkspacePlugin) RegisterRoutes(r *gin.Engine) {
	routes.SetUpWorkspaceRoutes(r)
}
