package payroll

import (
	"net/http"
	"payroll/api"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetJobDeapartments(c echo.Context) error {

	var jobDepartment api.JobDepartmentRequest
	if err := c.Bind(&jobDepartment); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.GetJobDepartments(jobDepartment)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) CreateJobDeapartment(c echo.Context) error {

	var jobDepartment api.JobDepartmentRequest

	if err := c.Bind(&jobDepartment); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.CreateDepartment(jobDepartment)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) GetJobDeapartmenById(c echo.Context) error {


	id := c.QueryParam("id")

	res, err := s.payroll.GetJobDepartmentById(id)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	

	return c.JSON(http.StatusOK, res)
}
