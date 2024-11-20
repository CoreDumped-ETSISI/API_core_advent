// middlewares/admin_access.go
package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type AdminAccess struct {
	IsAdmin     bool
	ExpiresAt   time.Time
	Attempts    int
	LastAttempt time.Time
}

var adminAccess = struct {
	sync.Mutex
	Access AdminAccess
}{}

// Load the admin password from the environment
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	adminPassword = os.Getenv("ADMIN_PASSWORD")
}

var adminPassword = os.Getenv("ADMIN_PASSWORD")

// AdminAccessMiddleware function
func AdminAccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminAccess.Lock()
		defer adminAccess.Unlock()

		// Check if admin access is still valid
		if adminAccess.Access.IsAdmin && time.Now().Before(adminAccess.Access.ExpiresAt) {
			c.Next() // Proceed to the admin route
			return
		}

		// Check time since last attempt
		if time.Since(adminAccess.Access.LastAttempt) > 60*time.Minute {
			adminAccess.Access.Attempts = 0 // Reset attempts after 10 minutes
		}

		// If attempts are exceeded
		if adminAccess.Access.Attempts >= 3 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many attempts. Please try again later."})
			c.Abort()
			return
		}

		// Check for the password in the body
		var json struct {
			Password string `json:"password"`
		}

		fmt.Println("json", json)
		if err := c.ShouldBindJSON(&json); err != nil || json.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password required"})
			c.Abort()
			return
		}

		// Check the password
		if json.Password != adminPassword {
			adminAccess.Access.Attempts++
			adminAccess.Access.LastAttempt = time.Now()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			c.Abort()
			return
		}

		// Grant admin access and set expiration
		adminAccess.Access.IsAdmin = true
		adminAccess.Access.ExpiresAt = time.Now().Add(5 * time.Minute) // Remember admin status for 5 minutes
		adminAccess.Access.Attempts = 0                                // Reset attempts on successful access

		c.JSON(http.StatusOK, gin.H{"message": "Admin access granted"})
	}
}
