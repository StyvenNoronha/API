package schema

import (
	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json: "active"`
}