package controllers

import (
	"document-management/models"
	"document-management/plugins/auth-plugin/services"
	"document-management/share/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

// contructor
func NewUserController(UserService *services.UserService) *UserController {
	return &UserController{UserService: UserService}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := uc.UserService.CreateUser(&user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Create Successfully", nil)
}

func (uc *UserController) Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Invalid request"})
		return
	}

	message, err := uc.UserService.Authenticate(loginRequest.Username, loginRequest.Password, c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, message, nil)
}
