package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Criando uma nova Instância do Echo
	e := echo.New()

	//Camada intermediária de comunicação
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// lista de Rotas
	e.GET("/students", getStudents)
	e.POST("/students", createStudents)

	e.GET("/students:id", getStudent)
	e.PUT("/students:id", updateStudent)
	e.DELETE("/students:id", deleteStudent)

	//Inicio do servidor
	e.Logger.Fatal(e.Start(":8080"))
}




// Gerenciar  a rota do servidor
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}

func createStudents(c echo.Context) error {
	return c.String(http.StatusOK, "Create students")
}

func getStudent(c echo.Context) error {
	id:= c.Param("id")
	return c.String(http.StatusOK, "Get student by id: "+id)
}

func updateStudent(c echo.Context) error {
	id:= c.Param("id")
	return c.String(http.StatusOK, "Update student by id: "+id)
}

func deleteStudent(c echo.Context) error {
	id:= c.Param("id")
	return c.String(http.StatusOK, "Deletar student by id: "+id)
}
