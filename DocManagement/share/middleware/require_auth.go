package middleware

import (
	"document-management/core/config"
	"document-management/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"time"
)

func RequireAuth(r *gin.Context) {

	fmt.Println("I'm in middleware")

	tokenString, err := r.Cookie("token")
	if err != nil {
		r.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			r.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("token expired")
		}

		var user *models.User
		config.ConnectDB().First(&user, claims["sub"])
		if user.Id == 0 {
			r.AbortWithStatus(http.StatusUnauthorized)
		}

		r.Set("user", user)
		r.Next()
		//fmt.Println(claims["foo"], claims["nbf"])
	} else {
		r.AbortWithStatus(http.StatusUnauthorized)
	}
}
