// controllers/problema_controller.go
package controllers

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/models"
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"slices"

)


func GetYear(c *gin.Context) {
	var years []int
	var problemas []models.Problema
	config.DB.Find(&problemas)
	for _, problema := range problemas {
		if slices.Contains(years, problema.Year) == false {
			years = append(years, problema.Year)
		}
	}
	c.JSON(http.StatusOK, years)

}


func GetProblema(c *gin.Context) {
    // Extract year and day from the URL parameters
    year := c.Param("year")
    day := c.Param("day")

    var problemas []models.Problema
    config.DB.Find(&problemas)

    // Load Spain's time zone
    timeZone, err := time.LoadLocation("Europe/Madrid")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to load time zone"})
        return
    }

    // Get current time in Spain's time zone
    currentTime := time.Now().In(timeZone)
    var results []string

    // Loop through the problemas and filter by year and day
    for _, problema := range problemas {
        // Check if the problema matches the year and day
        if fmt.Sprintf("%d", problema.Year) == year && fmt.Sprintf("%d", problema.Dia) == day {
            // Convert fecha_desbloqueo to Spain's time zone
            fechaDesbloqueo := problema.FechaDesbloqueo.In(timeZone)

            // Compare current time with fechaDesbloqueo
            if currentTime.After(fechaDesbloqueo) {
                results = append(results, problema.Enunciado) // Add the enunciado to the results
            } else {
                results = append(results, "No disponible todavía") // Not available yet
            }
        }
    }

    c.JSON(http.StatusOK, results)
}

func GetProblemas(c *gin.Context) {
    year := c.Param("year")
    var problemas []models.Problema
    config.DB.Where("year = ?", year).Find(&problemas)

    // Extract only "dia" into a slice of maps
    var problemasDias []map[string]int
    for _, problema := range problemas {
        problemasDias = append(problemasDias, map[string]int{"dia": problema.Dia})
    }

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
