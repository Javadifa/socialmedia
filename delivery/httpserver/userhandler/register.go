package userhandler

import (
	"net/http"

	"github.com/javadifa/socialmedia/param"
	"github.com/labstack/echo/v4"
)

func (h Handler) userRegister(c echo.Context) error {
	//a var
	var req param.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	//TODO : TECHNICAL DEBT: VALIDATOR AND FIELDSERROR

	resp, err := h.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}
