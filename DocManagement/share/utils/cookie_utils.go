package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetWorkspaceId get workspace id
func GetWorkspaceId(c *gin.Context) uint {
	workspaceIdStr, err := c.Cookie("workspaceId")
	if err != nil {
		return 0
	}
	workspaceId, err := strconv.ParseUint(workspaceIdStr, 10, 32)
	if err != nil {
		return 0
	}
	return uint(workspaceId)
}
