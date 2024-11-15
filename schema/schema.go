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

type StudentResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func NewResponse(students []Students) []StudentResponse {
	studentsResponse := []StudentResponse{}

	for _, student := range students{
		studentResponse := StudentResponse{
			ID: student.ID,
			Name: student.Name,
			CPF: student.CPF,
			Email: student.Email,
			Age: student.Age,
			Active: student.Active,
		}
		studentsResponse = append(studentsResponse,studentResponse)
	}
	return studentsResponse
}
