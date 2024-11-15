package db

import (
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct {
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
		log.Fatal().Err(err).Msgf("Failded to initialize SQLite: %s", err.Error())
	}

	//Gerenciar  as migrates
	db.AutoMigrate(&Students{})

	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

// Função para adicionar novo estudante
func (s *StudentHandler) AddStudent(student Students) error {
	result := s.DB.Create(&student)
	if result.Error != nil {
		log.Error().Msg("Error to create student")
		return result.Error
	} else {
		log.Info().Msg("Create student")
		return nil
	}

}
//Função para pegar todos os estudantes
func (s *StudentHandler) GetStudents() ([]Students, error) {
	students := []Students{}

	err := s.DB.Find(&students).Error
	return students, err
}

//Função para pegar um estudante especifico
func (s *StudentHandler) GetStudent(id int) (Students, error) {
	var student Students
	err := s.DB.First(&student, id)
	
	return student, err.Error
}

//Função para atualizar um estudante no banco de dados
func (s *StudentHandler) UpdateStudent(Updatestudent Students) error {
	return s.DB.Save(&Updatestudent).Error
}
