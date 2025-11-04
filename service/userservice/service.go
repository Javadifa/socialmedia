package userservice

import (
	"github.com/javadifa/socialmedia/entity"
)

//hash the password in service layer

//sanitizing input should happen in service layer

type Repository interface {
	IsEmailUnique(email string) (bool, error)
	IsHandleUnique(email string) (bool, error)
	Register(user entity.User) (entity.User, error)
	GetUserByHandle(handle string) (entity.User, bool, error)
	GetUserFeedByID(userID uint) ([]entity.Post, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}
type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(authGenerator AuthGenerator, repo Repository) Service {

	return Service{auth: authGenerator, repo: repo}
}
