package userservice

import (
	"fmt"

	"github.com/javadifa/socialmedia/entity"
	"github.com/javadifa/socialmedia/param"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {
	//2 ways:
	//pass : []byte(req.password, 0)
	//bcrypt.GenerateFromPassword(pass)

	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 0)
	//create new userservice in db
	user := entity.User{
		ID:        0,
		FullName:  req.FullName,
		Handle:    req.Handle,
		Email:     req.Email,
		Password:  string(hashedPwd),
		BirthDate: req.BirthDate,
		AvatarURL: req.AvatarURL,
	}

	//return the created userservice
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.RegisterResponse{
		FullName:  createdUser.FullName,
		Handle:    createdUser.Handle,
		Email:     createdUser.Email,
		BirthDate: createdUser.BirthDate,
		AvatarURL: createdUser.AvatarURL,
	}, nil

}
