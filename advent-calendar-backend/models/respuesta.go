// models/respuesta.go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Respuesta struct {
	gorm.Model
	UsuarioID        uint      `json:"usuario_id"`
	ProblemaID       uint      `json:"problema_id"`
	SolucionPropuesta string    `json:"solucion_propuesta"`
	FechaEnvio        time.Time `json:"fecha_envio"`
	Correcta          bool      `json:"correcta"`
}

func (Respuesta) TableName() string {
	return "Respuesta"
}