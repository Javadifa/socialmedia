package userhandler

import (
	"fmt"
	"net/http"

	"github.com/javadifa/socialmedia/param"
	"github.com/javadifa/socialmedia/service/authservice"
	"github.com/labstack/echo/v4"
)

func (h Handler) userFeed(c echo.Context) error {
	//authToken := c.Request().Header.Get("Authorization")
	//claims, err := h.authSvc.ParsToken(authToken)
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusUnauthorized)
	//}
	claims := c.Get("user")
	fmt.Println("claims:", claims)
	cl, ok := claims.(*authservice.Config)
	if !ok {
		fmt.Println("nok")
	} else {
		fmt.Println("cl:", cl)
	}
	
	resp, err := h.userSvc.GetUserFeed(param.FeedRequest{UserID: 0})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusOK, resp)
}
