package ms

import (
	"payroll/internal/config"
	httptransport "payroll/internal/controller"

	"github.com/labstack/echo/v4"
)

type Service struct {
	Echo   *echo.Echo
	HTTP   *httptransport.Service
	Config *config.Config
}
