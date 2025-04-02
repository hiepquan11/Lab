package repository

import (
	"document-management/models"
	"errors"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
	"path/filepath"
)

type DocumentRepository interface {
	UploadFile(file *multipart.FileHeader) (string, error)
	SaveDocument(doc *models.Document) error
	GetDocumentByUserId(userId uint) ([]*models.Document, error)
	DownloadFile(fileName string) (string, error)
}

type docRepo struct {
	DB *gorm.DB
}

// NewDocumentRepository constructor repository
func NewDocumentRepository(db *gorm.DB) DocumentRepository {
	return &docRepo{DB: db}
}

// UploadFile upload file
func (d docRepo) UploadFile(file *multipart.FileHeader) (string, error) {
	// implement
	uploadDir := "uploads"
	os.MkdirAll(uploadDir, os.ModePerm)

	filePath := filepath.Join(uploadDir, file.Filename)

	// open file upload
	content, err := file.Open()
	if err != nil {
		return "", err
	}
	defer content.Close()

	//create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// copy content from file upload into file destination
	if _, err := dst.ReadFrom(content); err != nil {
		return "", err
	}

	return filePath, nil
}

// SaveDocument save document name path into db
func (d docRepo) SaveDocument(doc *models.Document) error {
	// save document into db
	if err := d.DB.Save(doc).Error; err != nil {
		return err
	}
	return nil
}

// GetDocumentByUserId get list document
func (d docRepo) GetDocumentByUserId(userId uint) ([]*models.Document, error) {
	var documents []*models.Document
	if err := d.DB.Where("uploaded_by = ?", userId).Find(&documents).Error; err != nil {
		return nil, err
	}
	return documents, nil
}

// DownloadFile download repo
func (d docRepo) DownloadFile(fileName string) (string, error) {
	uploadDir := "uploads"
	filePath := filepath.Join(uploadDir, fileName)

	// check filepath is existed
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", errors.New("file not found")
		}
		return "", err
	}
	return filePath, nil
}
