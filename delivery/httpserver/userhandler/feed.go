package userhandler

import (
	"net/http"

	"github.com/javadifa/socialmedia/param"
	"github.com/labstack/echo/v4"
)

func (h Handler) userFeed(c echo.Context) error {
	//authToken := c.Request().Header.Get("Authorization")
	//claims, err := h.authSvc.ParsToken(authToken)
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusUnauthorized)
	//}
	claims := c.Get("user")

	resp, err := h.userSvc.GetUserFeed(param.FeedRequest{UserID: 0})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusOK, resp)
}
