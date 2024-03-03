package middleware

import (
	"github.com/labstack/echo/v4"
)

// emulator authen middleware
func IsAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("uid", "123")
		return next(c)
	}
}
