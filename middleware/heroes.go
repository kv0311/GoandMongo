package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LogInfo(e *echo.Echo) {
	e.Use((middleware.LoggerWithConfig(middleware.LoggerConfig{Format: "method=${method}, uri=${uri}, status=${status}\n"})))
}
func BasicAuthen(UserName, Password string, c echo.Context) (bool, error) {
	if UserName == "vinh" && Password == "123456" {
		return true, nil
	}
	return false, nil
}
