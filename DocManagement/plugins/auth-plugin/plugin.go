package auth_plugin

import (
	"document-management/plugins/auth-plugin/routes"
	"github.com/gin-gonic/gin"
)

type AuthPlugin struct{}

func NewAuthPlugin() *AuthPlugin {
	return &AuthPlugin{}
}

func (p *AuthPlugin) RegisterRoutes(r *gin.Engine) {
	routes.SetAuthUpRoutes(r)
}
