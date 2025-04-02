package controllers

import (
	"document-management/models"
	"document-management/plugins/document-plugin/services"
	"document-management/share/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DocumentController struct {
	docService *services.DocumentService
}

// NewDocumentController constructor controller
func NewDocumentController(docService *services.DocumentService) *DocumentController {
	return &DocumentController{docService: docService}
}

// UploadDocument upload doc
func (d *DocumentController) UploadDocument(r *gin.Context) {
	form, err := r.MultipartForm()
	if err != nil {
		utils.ErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	files := form.File["document"]
	if len(files) == 0 {
		utils.ErrorResponse(r, http.StatusBadRequest, "No files to upload")
		return
	}

	var docsUpload []*models.Document
	for _, file := range files {
		doc, err := d.docService.UploadFile(file, r)
		if err != nil {
			utils.ErrorResponse(r, http.StatusUnauthorized, err.Error())
			return
		}
		docsUpload = append(docsUpload, doc)
	}
	utils.SuccessResponse(r, http.StatusOK, "Upload Document Successfully", docsUpload)
}

func (d *DocumentController) GetDocumentByUserId(r *gin.Context) {
	userIdParams := r.Param("id")
	userId, err := strconv.ParseUint(userIdParams, 10, 64)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	documents, err := d.docService.GetDocumentByUserId(uint(userId))
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	r.JSON(http.StatusOK, gin.H{
		"message":   "Get successfully",
		"documents": documents,
	})
}

func (d *DocumentController) DownLoadController(r *gin.Context) {
	currentUser := utils.GetInfoUserFromToken(r)
	workspaceId := strconv.Itoa(int(utils.GetWorkspaceId(r)))

	allowed, err := utils.CheckPermission(workspaceId, strconv.Itoa(int(currentUser.Id)))
	if err != nil {
		utils.ErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	if !allowed {
		utils.ErrorResponse(r, http.StatusUnauthorized, "permission denied")
		return
	}

	var request struct {
		Filename string `json:"filename"`
	}

	// bind json from req body
	if err := r.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(r, http.StatusBadRequest, err.Error())
		return
	}

	filePath, err := d.docService.DownloadFile(request.Filename)
	if err != nil {
		utils.ErrorResponse(r, http.StatusInternalServerError, err.Error())
		return
	}
	r.File(filePath)

	//fileName := r.Param("filename")
	//
	//filePath, err := d.docService.DownloadFile(fileName)
	//if err != nil {
	//	r.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	//}
	//
	//r.Header("Content-Disposition", "attachment; filename="+fileName)
	//r.File(filePath)
}
