package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//handlers of nethttp
//func HealthCheck(writer http.ResponseWriter, request *http.Request) {}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is fine",
	})
}
