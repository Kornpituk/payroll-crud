package payroll

import (
	"net/http"
	"payroll/api"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetSalaries(c echo.Context) error {

	var salary api.SalaryRequest
	if err := c.Bind(&salary); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.GetSalaries(salary)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) CreateSalary(c echo.Context) error {

	var salary api.SalaryRequest

	if err := c.Bind(&salary); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.CreateSalary(salary)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) GetSalaryById(c echo.Context) error {


	id := c.QueryParam("id")

	res, err := s.payroll.GetSalaryById(id)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	

	return c.JSON(http.StatusOK, res)
}