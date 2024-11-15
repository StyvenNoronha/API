package api

import (
	"api/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_"api/docs"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

// @title Student API
// @version 1.0
// @description API para gerenciar alunos
// @host localhost:8080
// @BasePath /
// @schemes http
func NewServer() *API {
	// Criando uma nova Instância do Echo
	e := echo.New()

	//Camada intermediária de comunicação
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}

}

func (api *API) ConfigureRoutes() {
	// lista de Rotas
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudents)

	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
	//echoSwagger.WrapHandler
}

func (api *API) Start() error {

	//Inicio do servidor
	return api.Echo.Start(":8080")
}
