package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Usuario    string `gorm:"unique" json:"usuario"`
	Correo     string `gorm:"unique" json:"correo"`
	Contrasena string `json:"contrasena"`
	TiempoEspera int64  `json:"tiempo_espera"` // Tiempo de penalización en minutos
}

func (Usuario) TableName() string {
	return "Usuario"
}