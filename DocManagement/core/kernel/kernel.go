package kernel

import (
	authplugin "document-management/plugins/auth-plugin"
	docplugin "document-management/plugins/document-plugin"
	workspaceplugin "document-management/plugins/workspaces_plugin"
	"github.com/gin-gonic/gin"
)

type Plugin interface {
	RegisterRoutes(r *gin.Engine)
}

var plugins []Plugin

func InitKernel() {
	plugins = []Plugin{
		authplugin.NewAuthPlugin(),
		docplugin.NewDocumentPlugin(),
		workspaceplugin.NewDocumentPlugin(),
	}
}

func LoadPlugins(r *gin.Engine) {
	for _, plugin := range plugins {
		plugin.RegisterRoutes(r)
	}
}
