package payroll

import (
	"net/http"
	"payroll/api"

	"github.com/labstack/echo/v4"
)


func (s *Service) Createleave(c echo.Context) error {

	var leave api.LeaveRequest

	if err := c.Bind(&leave); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.CreateLeave(leave)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) Getleaves(c echo.Context) error {

	var leave api.LeaveRequest

	if err := c.Bind(&leave); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := s.payroll.GetLeaves(leave)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) GetleaveById(c echo.Context) error {


	id := c.QueryParam("id")

	res, err := s.payroll.GetLeaveById(id)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	

	return c.JSON(http.StatusOK, res)
}
