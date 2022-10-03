package ms

import (
	"net/http"
	"payroll/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func (s *Service) InitRoutes() {

	e := s.Echo
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	contextPath := s.Config.ContextPath

	e.GET(contextPath+"/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, domain.Health{
			Status: "OK",
		})
	})

	e.GET("/hello", s.HTTP.Payroll.Hello)


	e.GET("/payrolls", s.HTTP.Payroll.GetPayrolls)
	e.GET("/payroll/by", s.HTTP.Payroll.GetPayrollById)
	e.GET("/payroll/byEmp", s.HTTP.Payroll.GetPayrollByIdEmp)
	e.POST("/payroll", s.HTTP.Payroll.CreatePayroll)


	e.GET("/employees", s.HTTP.Payroll.GetEmployees)
	e.GET("/employee/by", s.HTTP.Payroll.GetEmployeeById)
	e.POST("/employee", s.HTTP.Payroll.CreateEmployee)

	e.GET("/salaries", s.HTTP.Payroll.GetSalaries)
	e.GET("/salary/by", s.HTTP.Payroll.GetSalaryById)
	e.POST("/salary", s.HTTP.Payroll.CreateSalary)

	e.POST("/leave", s.HTTP.Payroll.Createleave)
	e.GET("/leave/by", s.HTTP.Payroll.GetleaveById)
	e.GET("/leaves", s.HTTP.Payroll.Getleaves)

	e.POST("/jobdepartment", s.HTTP.Payroll.CreateJobDeapartment)
	e.GET("/jobdepartments", s.HTTP.Payroll.GetJobDeapartments)
	e.GET("/jobdepartment/by", s.HTTP.Payroll.GetJobDeapartmenById)



}
