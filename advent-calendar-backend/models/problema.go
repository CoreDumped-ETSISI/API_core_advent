// models/problema.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Problema struct {
	gorm.Model
	Year 		  int       `json:"year"`
	Dia           int       `json:"dia"`
	Titulo          string    `json:"titulo"`
	Enunciado       string    `json:"enunciado"`
	Solucion        string    `json:"solucion"`
	FechaDesbloqueo time.Time `json:"fecha_desbloqueo"`
	FechaBloqueo    time.Time `json:"fecha_bloqueo"`
}

func (Problema) TableName() string {
	return "Problema"
}
