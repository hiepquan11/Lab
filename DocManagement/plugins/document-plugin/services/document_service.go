package services

import (
	"document-management/models"
	"document-management/plugins/document-plugin/repository"
	"document-management/share/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"time"
)

type DocumentService struct {
	docRepo repository.DocumentRepository
}

// NewDocumentService constructor service
func NewDocumentService(docRepo repository.DocumentRepository) *DocumentService {
	return &DocumentService{docRepo: docRepo}
}

// UploadFile upload file
func (service *DocumentService) UploadFile(file *multipart.FileHeader, r *gin.Context) (*models.Document, error) {

	currentUser := utils.GetInfoUserFromToken(r)
	workspaceId := strconv.Itoa(int(utils.GetWorkspaceId(r)))

	// check permission
	allowed, err := utils.CheckPermission(workspaceId, strconv.Itoa(int(currentUser.Id)))
	if err != nil {
		return nil, err
	}
	//
	if !allowed {
		return nil, errors.New("permission denied")
	}

	// Upload file into database
	filePath, err := service.docRepo.UploadFile(file)
	if err != nil {
		return nil, err
	}

	// create model document
	doc := &models.Document{
		Name:       file.Filename,
		FilePath:   filePath,
		Workspace:  utils.GetWorkspaceId(r),
		UploadedAt: time.Now(),
		UploadedBy: currentUser.Id,
	}

	// save document into database
	if err := service.docRepo.SaveDocument(doc); err != nil {
		return nil, err
	}
	return doc, nil
}

// GetDocumentByUserId get doc by userid
func (service *DocumentService) GetDocumentByUserId(userID uint) ([]*models.Document, error) {
	return service.docRepo.GetDocumentByUserId(userID)
}

// DownloadFile download file
func (service *DocumentService) DownloadFile(fileName string) (string, error) {

	filePath, err := service.docRepo.DownloadFile(fileName)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
