// models/problema.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Problema struct {
	gorm.Model
	Titulo          string    `json:"titulo"`
	Enunciado       string    `json:"enunciado"`
	Solucion        string    `json:"solucion"`
	FechaDesbloqueo time.Time `json:"fecha_desbloqueo"`
}

func (Problema) TableName() string {
	return "Problema"
}