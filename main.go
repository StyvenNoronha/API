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

	// l ista de Rotas
	e.GET("/students", getStudents)

	//Inicio do servidor
	e.Logger.Fatal(e.Start(":8080"))
}

// Gerenciar  a rota do servidor
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}
