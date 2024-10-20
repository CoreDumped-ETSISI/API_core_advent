package config

import (
	"fmt"
	"log"
	"os"

	"advent-calendar-backend/models" // Importamos los modelos

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Crear la cadena DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	// Conectar a la base de datos
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Migrar las tablas
	err = DB.AutoMigrate(&models.Usuario{}, &models.Problema{}, &models.Respuesta{})
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}

	log.Println("Database connected and migrated successfully.")
}

// JWTSecret se toma de las variables de entorno
var JWTSecret = []byte(os.Getenv("JWT_SECRET"))
