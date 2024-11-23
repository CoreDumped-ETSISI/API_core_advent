package controllers

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/models"
	"advent-calendar-backend/utils"
	"net/http"
	"net/mail"

	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Valor      string `json:"Valor"`
	Contrasena string `json:"Contraseña"`
}

type RegisterInput struct {
	Correo     string `json:"Correo"`
	Usuario    string `json:"Usuario"`
	Contrasena string `json:"Contraseña"`
}

// Función de registro
func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		// imprimir error en consola
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Validar que el correo sea válido
	if !isValidEmail(input.Correo) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Correo no válido"})
		return
	}

	// Verificar si el usuario ya existe (por correo o usuario)
	var user models.Usuario
	//
	if err := config.DB.Where("correo = ? OR usuario = ?", input.Correo, input.Usuario).First(&user).Error; err == nil {
		// Devolver que el correo ya está en uso
		if user.Correo == input.Correo {
			c.JSON(http.StatusConflict, gin.H{"error": "El correo ya está en uso"})

		} else {
			c.JSON(http.StatusConflict, gin.H{"error": "El usuario ya está en uso"})
		}
		return
	}

	// Crear el nuevo usuario
	newUser := models.Usuario{
		Correo:     input.Correo,
		Contrasena: input.Contrasena,
		Usuario:    input.Usuario,
	}

	//fmt.Println(input.Contrasena)

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newUser})
}

// Validar correo con mail.ParseAddress
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Función de inicio de sesión
func Login(c *gin.Context) {
	var input LoginInput
	fmt.Println("Datos recibidos para registro:", input.Valor, input.Contrasena)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si el usuario existe (por correo o nombre de usuario)
	var user models.Usuario
	if err := config.DB.Where("correo = ? OR usuario = ?", input.Valor, input.Valor).First(&user).Error; err != nil {
		fmt.Println(err, "not found")

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Correo o usuario no encontrado"})
		return
	}

	fmt.Println(user.Contrasena == input.Contrasena)
	if user.Contrasena != input.Contrasena {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	// Generar el token
	token := utils.GenerateToken(user.ID)

	// Enviar el token como respuesta
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetRankingByYear(c *gin.Context) {
	year := c.Param("year")
	var ranking []struct {
		UsuarioID           uint    `json:"usuario_id"`
		Usuario             string  `json:"usuario"`
		Correctas           int     `json:"correctas"`
		TotalTimeDifference float64 `json:"total_time_difference"` // Total time difference in seconds
	}

	// Query to get the ranking
	err := config.DB.Table("Usuario").
		Select("Usuario.id as usuario_id, Usuario.usuario, COUNT(Respuesta.id) as correctas, SUM(julianday(Respuesta.fecha_envio) - julianday(Problema.fecha_desbloqueo)) as total_time_difference").
		Joins("LEFT JOIN Respuesta ON Respuesta.usuario_id = Usuario.id AND Respuesta.correcta = 1").
		Joins("LEFT JOIN Problema ON Respuesta.problema_id = Problema.id").
		Where("Respuesta.fecha_envio IS NOT NULL AND strftime('%Y', Respuesta.fecha_envio) = ? AND Respuesta.fecha_envio < Problema.fecha_bloqueo", year).
		Group("Usuario.id").
		Order("correctas DESC, total_time_difference ASC").
		Scan(&ranking).Error

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching ranking"})
		return
	}

	c.JSON(http.StatusOK, ranking)
}