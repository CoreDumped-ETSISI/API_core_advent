package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"advent-calendar-backend/config"
)

func GenerateToken(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(config.JWTSecret)
	return tokenString
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})
	return token, err
}
