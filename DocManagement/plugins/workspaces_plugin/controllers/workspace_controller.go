package controllers

import (
	"document-management/models"
	"document-management/plugins/workspaces_plugin/service"
	"document-management/share/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type WorkspaceController struct {
	workSpaceService *service.WorkspaceService
}

func NewWorkspaceController(workSpaceService *service.WorkspaceService) *WorkspaceController {
	return &WorkspaceController{workSpaceService: workSpaceService}
}

func (w *WorkspaceController) CreateWorkspace(r *gin.Context) {
	currentUser := utils.GetInfoUserFromToken(r)
	if currentUser == nil {
		utils.ErrorResponse(r, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var workSpaceRequest struct {
		Name string `json:"name"`
	}

	if err := r.ShouldBindJSON(&workSpaceRequest); err != nil {
		utils.ErrorResponse(r, http.StatusBadRequest, err.Error())
	}

	createWorkspace := &models.Workspace{
		Name:      workSpaceRequest.Name,
		OwnerId:   currentUser.Id,
		CreatedAt: time.Now(),
	}

	workspace, err := w.workSpaceService.CreateWorkspace(createWorkspace)
	if err != nil {
		utils.ErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}

	// assign permission
	utils.AssignWorkspacePermission(strconv.Itoa(int(currentUser.Id)), strconv.Itoa(int(workspace.Id)))

	r.SetCookie("workspaceId", strconv.Itoa(int(workspace.Id)), 3600, "/", "", false, true)
	utils.SuccessResponse(r, http.StatusCreated, "Create Successfully", workspace)

}
