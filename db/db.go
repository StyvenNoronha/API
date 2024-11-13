package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type Students struct{
	gorm.Model
	Name string `json:"name"`
	CPF int `json:"cpf"`
	Email string `json:"email"`
	Age int `json:"age"`
	Active bool `json: "active"`
}

  
 func Init() *gorm.DB{
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

 func AddStudent(student Students) error{
	db:= Init()
/*
	student := Students{
		Name: "Styven", ``
		CPF: 12332112336,
		Email: "styvenn16@gmail.com",
		Age:25,
		Active:true,
	}
*/
	result:= db.Create(&student)
	if result.Error != nil{
		return result.Error
	}else{
		
		fmt.Println("Deu tudo certo")
		return nil
	}
	
 }

func GetStudents()([]Students,error){
	students := []Students{}

	db:= Init()
	err:= db.Find(&students).Error
	return students,err
} 
  