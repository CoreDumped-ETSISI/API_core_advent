package controllers

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/models"
	"advent-calendar-backend/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type LoginInput struct {
	Valor    string `json:"Valor"`
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

	fmt.Println("Register", input)
	fmt.Println("Datos recibidos para registro:", input.Correo, input.Usuario, input.Contrasena)
	if err := c.ShouldBindJSON(&input); err != nil {
		// imprimir error en consola
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
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

// Función para obtener el ranking
func GetRanking(c *gin.Context) {
	var ranking []models.Usuario
	config.DB.Order("tiempo_espera asc").Find(&ranking)
	c.JSON(http.StatusOK, ranking)
}

// Función para obtener el ranking por año
func GetRankingByYear(c *gin.Context) {
	year := c.Param("year")
	var ranking []models.Usuario
	config.DB.Where("year = ?", year).Order("tiempo_espera asc").Find(&ranking)
	c.JSON(http.StatusOK, ranking)
}

