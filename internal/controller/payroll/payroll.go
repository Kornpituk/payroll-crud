package payroll

import (
	"net/http"
	"payroll/api"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetPayrolls(c echo.Context) error {

	var payroll api.PayrollRequest

	if err := c.Bind(&payroll); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.GetPayrolls(payroll)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) CreatePayroll(c echo.Context) error {

	var payroll api.PayrollRequest

	if err := c.Bind(&payroll); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.CreatePayroll(payroll)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) GetPayrollById(c echo.Context) error {


	id := c.QueryParam("id")

	res, err := s.payroll.GetPayrollById(id)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	

	return c.JSON(http.StatusOK, res)
}

func (s *Service) GetPayrollByIdEmp(c echo.Context) error {


	id := c.QueryParam("id")

	res, err := s.payroll.GetPayrollByIdEmp(id)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	

	return c.JSON(http.StatusOK, res)
}