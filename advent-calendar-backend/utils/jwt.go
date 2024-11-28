package utils

import (
	"advent-calendar-backend/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint, username string) string {
	// get username from the database User model
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"username":  username,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(config.JWTSecret)
	return tokenString
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorMalformed)
		}
		return config.JWTSecret, nil
	})

	if err != nil {
		return nil, err // Return error if token parsing fails
	}

	return token, nil // Return the token itself for further claims extraction
}
