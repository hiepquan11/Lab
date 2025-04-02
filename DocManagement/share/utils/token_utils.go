package utils

import (
	"document-management/models"
	"github.com/gin-gonic/gin"
)

// GetInfoUserFromToken Get user from token
func GetInfoUserFromToken(r *gin.Context) *models.User {
	user, exists := r.Get("user")
	if !exists {
		return nil
	}
	currentUser := user.(*models.User)
	return currentUser
}
