// controllers/respuesta_controller.go
package controllers

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/models"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func SubmitRespuesta(c *gin.Context) {
	var respuesta models.Respuesta
	if err := c.ShouldBindJSON(&respuesta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buscar al usuario
	var usuario models.Usuario
	if err := config.DB.First(&usuario, respuesta.UsuarioID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Verificar si el usuario tiene tiempo de espera
	if usuario.TiempoEspera > 0 {
		tiempoRestante := usuario.TiempoEspera - time.Now().Unix()
		if tiempoRestante > 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": "Debes esperar antes de enviar otra respuesta", "tiempo_restante": tiempoRestante})
			return
		}
	}

	// Verificar si la solución es correcta
	var problema models.Problema
	if err := config.DB.First(&problema, respuesta.ProblemaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problema no encontrado"})
		return
	}

	if respuesta.SolucionPropuesta == problema.Solucion {
		// Respuesta correcta, restablecer tiempo de espera
		usuario.TiempoEspera = 0
		respuesta.Correcta = true
		c.JSON(http.StatusOK, gin.H{"message": "¡Respuesta correcta!"})
	} else {
		// Respuesta incorrecta, duplicar penalización
		if usuario.TiempoEspera == 0 {
			usuario.TiempoEspera = time.Now().Unix() + 5*60 // Penalización inicial de 5 minutos
		} else {
			usuario.TiempoEspera = time.Now().Unix() + (usuario.TiempoEspera-time.Now().Unix())*2 // Duplicar el tiempo de espera
		}
		respuesta.Correcta = false
		c.JSON(http.StatusOK, gin.H{"message": "Respuesta incorrecta, penalización aplicada"})
	}

	// Guardar la respuesta y actualizar el usuario
	config.DB.Save(&respuesta)
	config.DB.Save(&usuario)
}