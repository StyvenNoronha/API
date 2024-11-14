package api

import (
	"api/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Gerenciar  a rota do servidor
// Função para listar todos os estudantes
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}
	return c.JSON(http.StatusOK, students)
}

// Função para cadastrar um novo estudante
func (api *API) createStudents(c echo.Context) error {
	student := db.Students{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student ")
	}

	return c.String(http.StatusOK, "Create students ")
}

// função para achar um determinado aluno
func (api *API) getStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Get student by id: "+id)
}

// Função para atualizar as informações de uma aluno
func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Update student by id: "+id)
}

// Função para deletar um aluno
func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Deletar student by id: "+id)
}