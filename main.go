package main

import (
	"fmt"
	"time"

	//"io"
	//"net/http"

	//"github.com/javadifa/social_media/repository/postgresql"
	//userservice "github.com/javadifa/social_media/service/userservice"
	//"github.com/labstack/echo/v4/middleware"

	"github.com/javadifa/socialmedia/config"
	"github.com/javadifa/socialmedia/delivery/httpserver"
	"github.com/javadifa/socialmedia/repository/postgresql"
	"github.com/javadifa/socialmedia/service/authservice"
	"github.com/javadifa/socialmedia/service/userservice"
)

// env-long
const (
	JwtSignKey                         = "jwt_secret"
	AccessTokenDuration  time.Duration = time.Hour * 24
	RefreshTokenDuration time.Duration = time.Hour * 24 * 30
	AccessTokenSubject                 = "access_token"
	RefreshTokenSubject                = "refresh_token"
)

func main() {
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenDuration,
			RefreshExpirationTime: RefreshTokenDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Postgresql: postgresql.Config{
			Username: "user",
			Password: "pass",
			Host:     "localhost",
			Port:     5433,
			DBName:   "mydb",
		},
	}

	authSvc, userSvc := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc)

	fmt.Println("start echo server")
	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)
	myPostgresqlRepo := postgresql.New(cfg.Postgresql)

	userSvc := userservice.New(authSvc, myPostgresqlRepo)

	return authSvc, userSvc

}
