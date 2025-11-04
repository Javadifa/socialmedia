package userhandler

import (
	"github.com/javadifa/socialmedia/service/authservice"
	"github.com/javadifa/socialmedia/service/userservice"
)

type Handler struct {
	authSvc     authservice.Service
	userSvc     userservice.Service
	authSignKey []byte
	//validator later on
}

func New(authSvc authservice.Service,
	userSvc userservice.Service, authSignKey string) Handler {
	return Handler{
		authSvc:     authSvc,
		userSvc:     userSvc,
		authSignKey: []byte(authSignKey),
	}
}
