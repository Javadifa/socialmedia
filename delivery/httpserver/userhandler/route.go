package userhandler

import (
	mw "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/user")

	//userGroup.GET("/profile", h.userProfile)
	//userGroup.GET("/feed", h.userFeed)

	userGroup.POST("/login", h.userLogin, mw.JWT(h.authSignKey))
	userGroup.POST("/register", h.userRegister)
	userGroup.GET("/feed", h.userFeed)

}
