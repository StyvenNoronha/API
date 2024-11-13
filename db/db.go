package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct{
	DB *gorm.DB
}

type Students struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json: "active"`
}

func Init() *gorm.DB {
	//Criando um banco com o gorm usando o Sqlite
	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	//Se der alguma erro entra no if
	if err != nil {
		log.Fatalln(err)
	}

	//Gerenciar  as migrates
	db.AutoMigrate(&Students{})

	return db
}

func NewStudentHandler(db *gorm.DB)*StudentHandler{
	return &StudentHandler{DB:db}
}

func (s *StudentHandler) AddStudent(student Students) error {
	result := s.DB.Create(&student)
	if result.Error != nil {
		return result.Error
	} else {

		fmt.Println("Deu tudo certo")
		return nil
	}

}

func (s *StudentHandler)  GetStudents() ([]Students, error) {
	students := []Students{}

	err := s.DB.Find(&students).Error
	return students, err
}
