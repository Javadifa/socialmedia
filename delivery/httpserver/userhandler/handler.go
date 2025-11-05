package userhandler

import (
	"github.com/javadifa/socialmedia/service/authservice"
	"github.com/javadifa/socialmedia/service/userservice"
)

type Handler struct {
	authSvc    authservice.Service
	userSvc    userservice.Service
	authConfig authservice.Config
	//validator later on
}

func New(authSvc authservice.Service,
	userSvc userservice.Service, authConfig authservice.Config) Handler {
	return Handler{
		authSvc:    authSvc,
		userSvc:    userSvc,
		authConfig: authConfig,
	}
}
