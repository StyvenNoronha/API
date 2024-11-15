package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool   `json: "active"` //*bool para aceitar nulo
}

func errParamRequired(param, tipo string) error {
	return fmt.Errorf("Param '%s' of type '%s' is required", param, tipo)
}

func (s *StudentRequest) Validate() error {
	if s.Name == "" {
		return errParamRequired("name", "string")
	}
	if s.CPF == 0 {
		return errParamRequired("cpf", "int")
	}
	if s.Email == "" {
		return errParamRequired("email", "string")
	}
	if s.Age == 0 {
		return errParamRequired("age", "int")
	}
	if s.Active ==  nil {
		return errParamRequired("active", "bool")
	}

	return nil
}