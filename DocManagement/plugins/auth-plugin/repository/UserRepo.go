package repository

import (
	"document-management/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// interface method
type UserRepository interface {
	CreateUser(user *models.User) error
	//GetUsers() ([]models.User, error)
	GetUserByUsername(name string) (*models.User, error)
}

// struct implement interface
type userRepo struct {
	DB *gorm.DB
}

// constructor repo
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{DB: db}
}

// implement CreateUser method
func (r *userRepo) CreateUser(user *models.User) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)
	if err != nil {
		return err
	}
	return r.DB.Create(user).Error
}

func (r *userRepo) GetUserByUsername(name string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", name).First(&user).Error
	return &user, err
}
