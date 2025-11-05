package userhandler

import (
	"github.com/javadifa/socialmedia/delivery/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/user")

	//userGroup.GET("/profile", h.userProfile)
	//userGroup.GET("/feed", h.userFeed)

	userGroup.POST("/login", h.userLogin, middleware.Auth(h.authSvc, h.authConfig))
	userGroup.POST("/register", h.userRegister)
	userGroup.GET("/feed", h.userFeed)

}
