package services

import (
	"document-management/models"
	"document-management/plugins/auth-plugin/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	// check user
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.New("username or password or email is null")
	}
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) Authenticate(name string, password string, r *gin.Context) (string, error) {
	user, err := s.UserRepo.GetUserByUsername(name)
	if err != nil {
		return "", errors.New("user not found")
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	GenerateToken(r, user.Id)

	return "Login Successfully", nil
}

// GenerateToken generate token
func GenerateToken(r *gin.Context, userId uint) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		r.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	r.SetSameSite(http.SameSiteLaxMode)
	r.SetCookie("token", tokenString, 15*60, "/", "", false, true)
}
