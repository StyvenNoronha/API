package api

import (
	"api/db"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() *API {
	// Criando uma nova Instância do Echo
	e := echo.New()

	//Camada intermediária de comunicação
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db := db.Init()
	return &API{
		Echo: e,
		DB:   db,
	}

}

func (api *API) ConfigureRoutes() {
	// lista de Rotas
	api.Echo.GET("/students", getStudents)
	api.Echo.POST("/students", createStudents)

	api.Echo.GET("/students:id", getStudent)
	api.Echo.PUT("/students:id", updateStudent)
	api.Echo.DELETE("/students:id", deleteStudent)
}

func (api *API) Start() error {

	//Inicio do servidor
	return api.Echo.Start(":8080")
}

// Gerenciar  a rota do servidor
// Função para listar todos os estudantes
func getStudents(c echo.Context) error {
	students, err := db.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}
	return c.JSON(http.StatusOK, students)
}

// Função para cadastrar um novo estudante
func createStudents(c echo.Context) error {
	student := db.Students{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := db.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student ")
	}

	return c.String(http.StatusOK, "Create students ")
}

// função para achar um determinado aluno
func getStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Get student by id: "+id)
}

// Função para atualizar as informações de uma aluno
func updateStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Update student by id: "+id)
}

// Função para deletar um aluno
func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Deletar student by id: "+id)
}
