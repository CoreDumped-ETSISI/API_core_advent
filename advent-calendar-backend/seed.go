package main

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/models"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func seed() {
	// Connect to the database
	config.ConnectDatabase()

	// Seed the database
	seedUsuarios()
	seedProblemas()
	seedRespuestas()

	log.Println("Database seeding completed successfully!")
}

func seedUsuarios() {
	usuarios := []models.Usuario{}

	// Generate 10 random users
	for i := 1; i <= 10; i++ {
		password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost) // Hash password
		usuarios = append(usuarios, models.Usuario{
			Usuario:     randomString(8),
			Correo:      randomString(5) + "@example.com",
			Contrasena:  string(password),
			TiempoEspera: 0,
		})
	}

	// Insert users into the database
	if err := config.DB.Create(&usuarios).Error; err != nil {
		log.Fatal("Failed to seed Usuarios: ", err)
	}
	log.Println("Usuarios seeded successfully!")
}


func seedProblemas() {
	problemas := []models.Problema{}
	now := time.Now()

	// Generate 10 random problems
	for i := 1; i <= 10; i++ {
		// Randomize the number of hours, minutes, and seconds to add to the current time
		randomHours := rand.Intn(24) + 1 // Random hours between 1 and 24 (1 day)
		randomMinutes := rand.Intn(60)   // Random minutes between 0 and 59
		randomSeconds := rand.Intn(60)   // Random seconds between 0 and 59

		problemas = append(problemas, models.Problema{
			Year:           now.Year(),
			Dia:            i,
			Titulo:         "Problema " + randomString(5),
			Enunciado:      "Enunciado del problema " + randomString(10),
			Solucion:       "Solucion del problema " + randomString(10),
			FechaDesbloqueo: now.Add(time.Hour * time.Duration(randomHours)).
				Add(time.Minute * time.Duration(randomMinutes)).
				Add(time.Second * time.Duration(randomSeconds)), // Add the random hours, minutes, and seconds
			FechaBloqueo: now.Add(time.Hour * 24 * time.Duration(i)), // Fixed as before
		})
	}

	// Insert problems into the database
	if err := config.DB.Create(&problemas).Error; err != nil {
		log.Fatal("Failed to seed Problemas: ", err)
	}
	log.Println("Problemas seeded successfully!")
}

func seedRespuestas() {
	respuestas := []models.Respuesta{}
	var usuarios []models.Usuario
	var problemas []models.Problema

	// Fetch all users and problems
	config.DB.Find(&usuarios)
	config.DB.Find(&problemas)

	// Generate random responses for each user-problem pair
	for _, usuario := range usuarios {
		for _, problema := range problemas {
			respuestas = append(respuestas, models.Respuesta{
				UsuarioID:        usuario.ID,
				ProblemaID:       problema.ID,
				SolucionPropuesta: "Respuesta de " + usuario.Usuario + " para " + problema.Titulo,
				FechaEnvio:        time.Now().Add(-time.Hour * time.Duration(rand.Intn(24))),
				Correcta:          rand.Intn(2) == 1, // Randomly mark as correct or incorrect
			})
		}
	}

	// Insert responses into the database
	if err := config.DB.Create(&respuestas).Error; err != nil {
		log.Fatal("Failed to seed Respuestas: ", err)
	}
	log.Println("Respuestas seeded successfully!")
}

// Utility function to generate a random string of given length
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
