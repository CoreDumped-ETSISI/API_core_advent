// controllers/respuesta_controller.go
package controllers

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/models"
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)


func getUserIDFromToken(c *gin.Context) (uint, error) {
	if userID, exists := c.Get("user_id"); exists {
		return userID.(uint), nil // Type assertion to uint
	}
	return 0, fmt.Errorf("user ID not found in context")
}

func SubmitRespuesta(c *gin.Context) {
	var respuesta models.Respuesta
	if err := c.ShouldBindJSON(&respuesta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from the token (assumes you have a middleware that handles this)
	userID, err := getUserIDFromToken(c) // You need to implement this function
	fmt.Println("userID", userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get year and day from URL parameters
	year := c.Param("year")
	day := c.Param("day")

	// Find the problem using year and day
	var problema models.Problema
	if err := config.DB.Where("year = ? AND dia = ?", year, day).First(&problema).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problema no encontrado"})
		return
	}

	// Set the UserID from the token
	respuesta.UsuarioID = userID

	// Find the user based on userID
	var usuario models.Usuario
	if err := config.DB.First(&usuario, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Check if the user has a waiting time
	if usuario.TiempoEspera > 0 {
		tiempoRestante := usuario.TiempoEspera - time.Now().Unix()
		if tiempoRestante > 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": "Debes esperar antes de enviar otra respuesta", "tiempo_restante": tiempoRestante})
			return
		}
	}

	// Check if the proposed solution is correct
	if respuesta.SolucionPropuesta == problema.Solucion {
		// Correct answer, reset waiting time
		usuario.TiempoEspera = 0
		respuesta.Correcta = true
		c.JSON(http.StatusOK, gin.H{"message": "¡Respuesta correcta!"})
	} else {
		// Incorrect answer, double the penalty
		if usuario.TiempoEspera == 0 {
			usuario.TiempoEspera = time.Now().Unix() + 5*60 // Initial 5 minutes penalty
		} else {
			usuario.TiempoEspera = time.Now().Unix() + (usuario.TiempoEspera-time.Now().Unix())*2 // Double the waiting time
		}
		respuesta.Correcta = false
		c.JSON(http.StatusOK, gin.H{"message": "Respuesta incorrecta, penalización aplicada"})
	}

	// Save the response and update the user
	if err := config.DB.Create(&respuesta).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la respuesta"})
		return
	}
	if err := config.DB.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
		return
	}
}




func GetRespuestasByUsuario(c *gin.Context) {
	id := c.Param("id")
	var respuestas []models.Respuesta
	config.DB.Where("usuario_id = ?", id).Find(&respuestas)
	c.JSON(http.StatusOK, respuestas)
}

func GetRespuestasByProblema(c *gin.Context) {
	id := c.Param("id")
	var respuestas []models.Respuesta
	config.DB.Where("problema_id = ?", id).Find(&respuestas)
	c.JSON(http.StatusOK, respuestas)
}

func GetRespuestasByUsuarioAndProblema(c *gin.Context) {

	usuarioID := c.Param("usuario_id")
	problemaID := c.Param("problema_id")

	var respuestas []models.Respuesta
	config.DB.Where("usuario_id = ? AND problema_id = ?", usuarioID, problemaID).Find(&respuestas)
	c.JSON(http.StatusOK, respuestas)
}

func GetRespuestasByUsuarioAndCorrecta(c *gin.Context) {
	usuarioID := c.Param("usuario_id")
	correcta := c.Param("correcta")

	var respuestas []models.Respuesta
	config.DB.Where("usuario_id = ? AND correcta = ?", usuarioID, correcta).Find(&respuestas)
	c.JSON(http.StatusOK, respuestas)
}

func GetRespuestasByProblemaAndCorrecta(c *gin.Context) {
	problemaID := c.Param("problema_id")
	correcta := c.Param("correcta")

	var respuestas []models.Respuesta
	config.DB.Where("problema_id = ? AND correcta = ?", problemaID, correcta).Find(&respuestas)
	c.JSON(http.StatusOK, respuestas)
}

