package main

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"

	"time"
	_ "time/tzdata"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}
	loc, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		log.Println("Error loading timezone")
	}
	time.Local = loc // -> this is setting the global timezone

	// Conectar la base de datos
	config.ConnectDatabase()
	gin.SetMode(gin.ReleaseMode)

	// Iniciar el servidor
	r := routes.SetupRouter()
	r.SetTrustedProxies([]string{""})
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Permitir todos los orígenes
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Puerto por defecto si no se define en las variables de entorno
	}
	r.Run(":" + port)
}
