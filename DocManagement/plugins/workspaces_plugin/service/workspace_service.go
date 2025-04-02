package service

import (
	"document-management/models"
	"document-management/plugins/workspaces_plugin/repository"
)

type WorkspaceService struct {
	workSpaceRepo repository.WorkSpacesRepo
}

func NewWorkspaceService(workSpaceRepo repository.WorkSpacesRepo) *WorkspaceService {
	return &WorkspaceService{workSpaceRepo: workSpaceRepo}
}

func (w *WorkspaceService) CreateWorkspace(wp *models.Workspace) (*models.Workspace, error) {
	return w.workSpaceRepo.CreateWorkspace(wp)
}
