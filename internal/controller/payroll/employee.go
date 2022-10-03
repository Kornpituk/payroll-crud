package payroll

import (
	"net/http"
	"payroll/api"

	"github.com/labstack/echo/v4"
)


func (s *Service) CreateEmployee(c echo.Context) error {

	var employee api.EmployeeRequest

	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.CreateEmployee(employee)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) GetEmployees(c echo.Context) error {

	var employee api.EmployeeRequest

	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.GetEmployees(employee)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) GetEmployeeById(c echo.Context) error {


	var sal api.SalaryRequest
	id := c.QueryParam("id")

	res, err := s.payroll.GetEmployeeById(id,sal)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	

	return c.JSON(http.StatusOK, res)
}
