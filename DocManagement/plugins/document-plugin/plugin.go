package document_plugin

import (
	"document-management/plugins/document-plugin/routes"
	"github.com/gin-gonic/gin"
)

type DocumentPlugin struct{}

func NewDocumentPlugin() *DocumentPlugin {
	return &DocumentPlugin{}
}

func (p *DocumentPlugin) RegisterRoutes(r *gin.Engine) {
	routes.SetDocumentUpRoutes(r)
}
