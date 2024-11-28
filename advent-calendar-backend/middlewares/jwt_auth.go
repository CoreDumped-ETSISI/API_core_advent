package middlewares

import (
	"advent-calendar-backend/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token del header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		// Comprobar que el formato sea "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato del token no válido"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no válido o expirado"})
			c.Abort()
			return
		}

		// Pasar el token decodificado al contexto
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID, idOk := claims["user_id"].(string)
			username, usernameOk := claims["username"].(string)

			if !idOk || !usernameOk {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Reclamaciones del token no válidas"})
				c.Abort()
				return
			}

			// Establecer valores en el contexto
			c.Set("user_id", userID)
			c.Set("username", username)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Reclamaciones no válidas"})
			c.Abort()
			return
		}

		c.Next()
	}
}
