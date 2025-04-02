package repository

import (
	"document-management/models"
	"gorm.io/gorm"
)

type WorkSpacesRepo interface {
	CreateWorkspace(workSpace *models.Workspace) (*models.Workspace, error)
}

type workSpaceRepo struct {
	DB *gorm.DB
}

func NewWorkSpacesRepo(db *gorm.DB) WorkSpacesRepo {
	return &workSpaceRepo{DB: db}
}

func (w *workSpaceRepo) CreateWorkspace(workSpace *models.Workspace) (*models.Workspace, error) {
	result := w.DB.Create(workSpace)
	if result.Error != nil {
		return nil, result.Error
	}
	return workSpace, nil
}
