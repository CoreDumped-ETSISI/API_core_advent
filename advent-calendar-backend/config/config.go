package config

import (
	"advent-calendar-backend/models" // Importamos los modelos
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Load environment variables from .env
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	// Check if required variables are set
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPath := os.Getenv("DB_PATH")

	if dbPassword == "" || dbPath == "" {
		log.Fatal("DB_PASSWORD or DB_PATH environment variable is required")
	}
	// Define the SQLite database file path with encryption
	dbFilePath := fmt.Sprintf("file:%s?_pragma_key=%s", dbPath, dbPassword)

	// Connect to the SQLite database with encryption
	DB, err = gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to SQLite database with encryption!", err)
	}

	// Migrate the tables
	err = DB.AutoMigrate(&models.Usuario{}, &models.Problema{}, &models.Respuesta{})
	if err != nil {
		log.Fatal("Failed to migrate SQLite database!", err)
	}

	log.Println("Encrypted SQLite database connected and migrated successfully.")
}

// JWTSecret se toma de las variables de entorno
var JWTSecret = []byte(os.Getenv("JWT_SECRET"))
