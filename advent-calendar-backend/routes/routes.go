package routes

import (
    "advent-calendar-backend/controllers"
    "advent-calendar-backend/middlewares"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,

	}
	r.Use(cors.New(corsConfig))


    r.GET("/", controllers.GetYear)

    // Public routes
    r.POST("/login", controllers.Login)
    r.POST("/register", controllers.Register)

    // Admin routes
    adminRoutes := r.Group("/admin")

	adminRoutes.POST("/", middlewares.AdminAccessMiddleware())

    adminRoutes.Use(middlewares.AdminAccessMiddleware())
    {
        adminRoutes.POST("/problemas", controllers.CreateProblema)
        adminRoutes.PUT("/problemas/:year/:id", controllers.UpdateProblema)
        adminRoutes.DELETE("/problemas/:year/:id", controllers.DeleteProblema)
    }

    r.GET("/:year/", controllers.GetProblemas)
    r.GET("/ranking/:year", controllers.GetRankingByYear)
    r.GET("/:year/resueltas/:usuario", controllers.GetRespuestasByUsuarioAndCorrecta)
	r.Use(middlewares.JWTAuthMiddleware())
	{
		r.GET("/:year/:day", controllers.GetProblema)
		r.POST("/:year/:day", controllers.SubmitRespuesta)
	}


    return r
}
