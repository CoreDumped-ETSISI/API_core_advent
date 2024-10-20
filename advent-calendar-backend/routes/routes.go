package routes

import (
	"advent-calendar-backend/controllers"
	"advent-calendar-backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

)

func SetupRouter() *gin.Engine {
	// GIN_MODE RELLEASE
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // Permitir todos los orígenes
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	


    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Advent Calendar API!",
        })
    })

	// Rutas públicas

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	// Rutas protegidas
	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuthMiddleware())  // Middleware JWT
	{
		protected.GET("/problemas", controllers.GetProblemas)
		protected.POST("/respuestas", controllers.SubmitRespuesta)
	}

	return r
}
