package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWTToken(user_id uint, user_role string) (string, error) {
	if user_role == "" {
		user_role = "basic"
	}

	claims := jwt.MapClaims{
		"authenticated": true,
		"user_id":       user_id,
		"user_role":     user_role,
		"exp":           time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":           time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
