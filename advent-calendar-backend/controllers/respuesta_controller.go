// controllers/respuesta_controller.go
package controllers

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/models"
	"advent-calendar-backend/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


type RespuestaInput struct {
	SolucionPropuesta string `json:"solucion_propuesta"`
}


func SubmitRespuesta(c *gin.Context) {
	var input RespuestaInput // Use RespuestaInput for binding
	var respuesta models.Respuesta

	// Bind incoming JSON to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID from the token with ValidateToken function
	token, err := utils.ValidateToken(c.GetHeader("Authorization")[7:])
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort() // Stop further processing
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	// Get year and day from URL parameters
	year := c.Param("year")
	day := c.Param("day")

	// Find the problem using year and day
	var problema models.Problema
	if err := config.DB.Where("year = ? AND dia = ?", year, day).First(&problema).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Problema no encontrado"})
		return
	}


    // Check if the user has already submitted a correct answer for this problem
    var existingRespuesta models.Respuesta

    if err := config.DB.Where("usuario_id = ? AND problema_id = ? AND correcta = true", userID, problema.ID).First(&existingRespuesta).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Ya has enviado una respuesta correcta para este problema"})
        return
    }

	// Set the UserID from the token
	respuesta.UsuarioID = userID
	respuesta.FechaEnvio = time.Now()
	respuesta.SolucionPropuesta = input.SolucionPropuesta // Use input for the solution proposed
	respuesta.ProblemaID = problema.ID
	

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

	usuario := c.Param("usuario")
	problemaID := c.Param("problema_id")

	// sacamos usuarioID
	var usuarioID uint
	config.DB.Where("usuario = ? or correo = ?", usuario, usuario).First(&usuarioID)


	var respuestas []models.Respuesta
	config.DB.Where("usuario_id = ? AND problema_id = ?", usuarioID, problemaID).Find(&respuestas)
	c.JSON(http.StatusOK, respuestas)
}

func GetRespuestasByUsuarioAndCorrecta(c *gin.Context) {

	// Get the user ID from the token with ValidateToken function
	token, err := utils.ValidateToken(c.GetHeader("Authorization")[7:])
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort() // Stop further processing
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	usuarioID := uint(claims["user_id"].(float64))
	// pasamos year a int
	yearParam := c.Param("year")
	year, err := strconv.Atoi(yearParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year format"})
		return
	}
	
	var respuestas []models.Respuesta
	var esteuser []models.Respuesta

	problemasIDvalidos := []uint{}

	// sacamos de la tabla problemas los problemas que sean del año
	var problemas []models.Problema
	config.DB.Where("year = ?", year).Find(&problemas)
	for _, problema := range problemas {
		problemasIDvalidos = append(problemasIDvalidos, problema.ID)
	}

	// sacamos de la tabla respuestas las respuestas que sean correctas y del usuario y los id de los problemas pertenecientes al año
	config.DB.Where("usuario_id = ? AND correcta = ?", usuarioID, true).Find(&esteuser)

	for _, respuesta := range esteuser {
		for _, problemaID := range problemasIDvalidos {
			if respuesta.ProblemaID == problemaID {
				respuestas = append(respuestas, respuesta)
			}
		}
	}

	// si es nulo devolvemos un mensaje
	if len(respuestas) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay respuestas correctas para el año seleccionado"})
		return 
	}


	// devolvemos el dia de cada problema y el titulo de dicho problema en json
	var respuestasJSON []gin.H

	for _, respuesta := range respuestas {
		var problema models.Problema
		config.DB.Where("id = ?", respuesta.ProblemaID).First(&problema)
		respuestasJSON = append(respuestasJSON, gin.H{
			"dia":    problema.Dia,
			"titulo": problema.Titulo,
		})
	}


	c.JSON(http.StatusOK, respuestasJSON)
}

func GetRespuestasByProblemaAndCorrecta(c *gin.Context) {

	problemaID := c.Param("problema_id")
	correcta := c.Param("correcta")
	var respuestas []models.Respuesta
	config.DB.Where("problema_id = ? AND correcta = ?", problemaID, correcta).Find(&respuestas)
	c.JSON(http.StatusOK, respuestas)
}

