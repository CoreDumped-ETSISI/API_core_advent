package main

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/routes"
	"log"
	"os"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	// Conectar la base de datos
	config.ConnectDatabase()

	// Iniciar el servidor
	r := routes.SetupRouter()
	r.SetTrustedProxies([]string{""})
	r.Use(cors.New(cors.Config{
        AllowAllOrigins: true, // Permitir todos los orígenes
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	gin.SetMode(gin.ReleaseMode)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Puerto por defecto si no se define en las variables de entorno
	}
	r.Run(":" + port)
}
