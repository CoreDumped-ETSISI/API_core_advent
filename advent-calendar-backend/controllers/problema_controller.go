// controllers/problema_controller.go
package controllers

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"slices"
    "sort"
	"time"
)


func GetYear(c *gin.Context) {
    var years []int
    var problemas []models.Problema
    config.DB.Find(&problemas)

    for _, problema := range problemas {
        if !slices.Contains(years, problema.Year) {
            years = append(years, problema.Year)
        }
    }

    // Ordenar la lista de años
    sort.Ints(years)

    c.JSON(http.StatusOK, years)
}

func GetProblema(c *gin.Context) {
    // Extract user ID from JWT stored in the context by JWTAuthMiddleware
    claims, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo autenticar el usuario"})
        return
    }

    // Convert user ID from the claims, dchange it to string
    var usuarioID string
    switch v := claims.(type) {
    case string:
        usuarioID = v
    case float64:
        usuarioID = fmt.Sprintf("%.0f", v) // Convert float64 to string
    default:
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Tipo de usuario no válido"})
        return
    }

    dia := c.Param("day")
	year := c.Param("year")

    // Retrieve the problem using the problemaID
    var problema models.Problema
    if err := config.DB.First(&problema, "year = ? AND dia = ?", year, dia).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Problema no encontrado"})
		return
	}
	
    // Check if the problem is unlocked (FechaDesbloqueo has passed)
    currentTime := time.Now()
    if currentTime.Before(problema.FechaDesbloqueo) {
        // Problem is not unlocked yet, return time left to unlock
        timeLeft := problema.FechaDesbloqueo.Sub(currentTime).String()
        c.JSON(http.StatusOK, gin.H{
            "desbloqueado":      false,
            "tiempo_para_desbloquear": timeLeft,
        })
        return
    }

	problemaID := problema.ID
    // Problem is unlocked, return the problem and answer status
    var respuesta models.Respuesta
    if err := config.DB.Where("usuario_id = ? AND problema_id = ? AND correcta = 1", usuarioID, problemaID).First(&respuesta).Error; err == nil {
        // User has already submitted a valid answer
        c.JSON(http.StatusOK, gin.H{
            "problema":          problema.Enunciado,
            "respuesta_valida":  true,
            "respuesta_usuario": respuesta,
        })
        return
    }

    // User has not submitted a valid answer yet
    c.JSON(http.StatusOK, gin.H{
        "problema":         problema.Enunciado,
        "respuesta_valida": false,
    })
}

func GetProblemas(c *gin.Context) {
    year := c.Param("year")
    var problemas []models.Problema
    config.DB.Where("year = ?", year).Find(&problemas)

    // Crear una lista de enteros con los días
    var problemasDias []int
    for _, problema := range problemas {
        problemasDias = append(problemasDias, problema.Dia)
    }

    // Ordenar la lista de días
    sort.Ints(problemasDias)

    c.JSON(http.StatusOK, problemasDias)
}


func CreateProblema(c *gin.Context) {
	var input models.Problema

	if err := c.ShouldBindJSON(&input); err != nil {
		// imprimir error en consola
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating problema"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdateProblema(c *gin.Context) {
	var problema models.Problema
	if err := c.ShouldBindJSON(&problema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&problema)
	c.JSON(http.StatusOK, problema)
}

func DeleteProblema(c *gin.Context) {
	id := c.Param("id")
	config.DB.Where("id = ?", id).Delete(&models.Problema{})
	c.JSON(http.StatusOK, gin.H{"data": id})
}


func GetRespuestas(c *gin.Context) {
	var respuestas []models.Respuesta
	if err := config.DB.Find(&respuestas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve respuestas"})
		return
	}
	c.JSON(http.StatusOK, respuestas)
}

func CreateRespuesta(c *gin.Context) {
	var respuesta models.Respuesta
	if err := c.ShouldBindJSON(&respuesta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&respuesta)
	c.JSON(http.StatusOK, respuesta)
}

func UpdateRespuesta(c *gin.Context) {
	var respuesta models.Respuesta
	if err := c.ShouldBindJSON(&respuesta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&respuesta)
	c.JSON(http.StatusOK, respuesta)
}

func DeleteRespuesta(c *gin.Context) {
	id := c.Param("id")
	config.DB.Where("id = ?", id).Delete(&models.Respuesta{})
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func GetInfoProblemas(c *gin.Context) {
    var problemas []models.Problema
    config.DB.Find(&problemas)
    c.JSON(http.StatusOK, problemas)
}
