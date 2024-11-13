package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type Students struct{
	gorm.Model
	Name string
	CPF int
	Email string
	Age int
	Active bool
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

 func AddStudent(){
	db:= Init()

	student := Students{
		Name: "Styven",
		CPF: 12332112336,
		Email: "styvenn16@gmail.com",
		Age:25,
		Active:true,
	}

	result:= db.Create(&student)
	if result.Error != nil{
		fmt.Println("Deu erro")
	}else{
		fmt.Println("Deu tudo certo")
	}
 }
  