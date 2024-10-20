// controllers/problema_controller.go
package controllers

import (
	"advent-calendar-backend/config"
	"advent-calendar-backend/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetProblemas(c *gin.Context) {
	var problemas []models.Problema
	config.DB.Find(&problemas)
	c.JSON(http.StatusOK, problemas)
}
