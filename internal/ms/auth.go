package ms

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

var (
	jwtSecret = []byte("secret")
	expiresAt = time.Minute * 60
)

func createToken() (string, error) {
	claims := &jwtCustomClaims{
		"Super Admin",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresAt).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// TODO: Implement authentication logic
	if username != "admin" || password != "admin" {
		return echo.ErrUnauthorized
	}

	// TODO: Pass user object to createToken with claims
	t, err := createToken()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
