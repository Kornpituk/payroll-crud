package payroll
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) Hello(c echo.Context) error {
	res, err := s.payroll.Hello()
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}
	return c.JSON(http.StatusOK, res)
}
