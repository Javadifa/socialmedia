package httpserver

import (
	"fmt"

	"github.com/javadifa/socialmedia/config"
	"github.com/javadifa/socialmedia/delivery/httpserver/userhandler"
	"github.com/javadifa/socialmedia/service/authservice"
	"github.com/javadifa/socialmedia/service/userservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service) Server {
	return Server{
		config:      config,
		userHandler: userhandler.New(authSvc, userSvc, config.Auth.SignKey),
	}
}

func (s Server) Serve() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health-check", s.healthCheck)

	s.userHandler.SetUserRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
