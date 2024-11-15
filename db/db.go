package db

import (
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"api/schema"
	
)

type StudentHandler struct {
	DB *gorm.DB
}



func Init() *gorm.DB {
	//Criando um banco com o gorm usando o Sqlite
	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	//Verificando se o banco foi criado
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize SQLite: %s", err.Error())
	}

	//Gerenciar  as migrates
	db.AutoMigrate(&schema.Students{})

	return db
}
// Função para criar um novo estudante
func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

// Função para adicionar novo estudante
func (s *StudentHandler) AddStudent(student schema.Students) error {
	//Criando um novo estudante
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
func (s *StudentHandler) GetStudents() ([]schema.Students, error) {
	students := []schema.Students{}

	err := s.DB.Find(&students).Error
	return students, err
}

//Função para pegar um estudante especifico
func (s *StudentHandler) GetStudent(id int) (schema.Students, error) {
	var student schema.Students
	err := s.DB.First(&student, id)
	
	return student, err.Error
}

//Função para atualizar um estudante no banco de dados
func (s *StudentHandler) UpdateStudent(Updatestudent schema.Students) error {
	return s.DB.Save(&Updatestudent).Error
}

//Função para deletar um estudante no banco de dados
func (s *StudentHandler) DeleteStudent(student schema.Students) error {
	return s.DB.Delete(&student).Error
}