package api

import (
	"api/schema"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	
)

// Gerenciar  a rota do servidor
// Função para listar todos os estudantes

// getStudents godoc
// @Summary Listar todos os estudantes
// @Description Retorna uma lista de todos os estudantes
// @Tags students
// @Accept json
// @Produce json
// @Param active query string false "Filtrar por ativo"
// @Success 200 {object} schema.StudentResponse
// @Router /students [get]
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	//lista de estudantes ativos ou não
	active := c.QueryParam("active")
	if active != "" {
		ativo, err := strconv.ParseBool(active)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid active value")
		}
		//método para filtrar os estudantes ativos ou não
		students, err = api.DB.GetActiveStudents(ativo)

	}

	listStudents := map[string][]schema.StudentResponse{"students": schema.NewResponse(students)}

	return c.JSON(http.StatusOK, listStudents)
}

// Função para cadastrar um novo estudante
// @Summary Cadastrar um novo estudante
// @Description Cadastra um novo estudante
// @Tags students
// @Accept json
// @Produce json
// @Sucess 200 {object} schema.StudentResponse
// @Router /students [post]
func (api *API) createStudents(c echo.Context) error {
	studentReq := StudentRequest{}
	if err := c.Bind(&studentReq); err != nil {
		return err
	}

	if err := studentReq.Validate(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	student := schema.Students{
		Name:   studentReq.Name,
		CPF:    studentReq.CPF,
		Email:  studentReq.Email,
		Age:    studentReq.Age,
		Active: *studentReq.Active,
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student ")
	}

	return c.JSON(http.StatusOK, student)
}

// função para achar um determinado aluno
// getStudent godoc
// @Summary Buscar um estudante
// @Description Retorna um estudante
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {object} schema.StudentResponse
// @Router /students/{id} [get]
func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get student id")
	}
	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusNotFound, "server not found")
	}
	return c.JSON(http.StatusOK, student)
}

// Função para atualizar as informações de uma aluno
// updateStudent godoc
// @Summary Atualizar um estudante
// @Description Atualiza um estudante
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {object} schema.StudentResponse
// @Router /students/{id} [put]
func (api *API) updateStudent(c echo.Context) error {
	// ela transforma o id em um inteiro
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get student id")
	}
	// Criando um novo estudante
	receiveStudent := schema.Students{}
	if err := c.Bind(&receiveStudent); err != nil {
		return err
	}

	// Atualizando o estudante
	updateStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusNotFound, "server not found")
	}
	/// Atualizando o estudante
	student := updateStudentInfo(receiveStudent, updateStudent)

	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to update student")
	}

	return c.JSON(http.StatusOK, student)
}

// Função para atualizar as informações de uma aluno
func updateStudentInfo(receiveStudent, updateStudent schema.Students) schema.Students {
	// Atualizando o estudante
	if receiveStudent.Name != "" {
		updateStudent.Name = receiveStudent.Name
	}
	if receiveStudent.Email != "" {
		updateStudent.Email = receiveStudent.Email
	}
	if receiveStudent.CPF > 0 {
		updateStudent.CPF = receiveStudent.CPF
	}
	if receiveStudent.Age > 0 {
		updateStudent.Age = receiveStudent.Age
	}
	if receiveStudent.Active != updateStudent.Active {
		updateStudent.Active = receiveStudent.Active
	}
	return updateStudent
}

// Função para deletar um aluno
// deleteStudent godoc
// @Summary Deletar um estudante
// @Description Deleta um estudante
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {object} schema.StudentResponse
// @Router /students/{id} [delete]
func (api *API) deleteStudent(c echo.Context) error {
	// ela transforma o id em um inteiro
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get student id")
	}
	// Criando um novo estudante
	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusNotFound, "server not found")
	}
	/// Deletando o estudante
	if err := api.DB.DeleteStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete student")
	}

	return c.JSON(http.StatusOK, student)
}
