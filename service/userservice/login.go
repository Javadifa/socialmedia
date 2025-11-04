package userservice

import (
	"fmt"

	"github.com/javadifa/socialmedia/param"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//TODO: based on solid maybe it was better to check the existence and getting
//todo: return a complete response containing info and tokens
//in our case we must separate due ro search in database to be shown on search bar

func (s Service) Login(req param.LoginRequest) (param.Tokens, error) {

	user, exist, repoErr := s.repo.GetUserByHandle(req.Handle)
	if repoErr != nil {
		return param.Tokens{}, fmt.Errorf("couldn't fetch the userservice by handle: %w", repoErr)
	}

	if !exist {
		return param.Tokens{}, fmt.Errorf("sth went wrong with either handle or password")
	}

	//todo: cost > 13
	err := bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		return param.Tokens{}, fmt.Errorf("sth went wrong with either handle or password")
	}

	accessToken, jErr := s.auth.CreateAccessToken(user)
	if jErr != nil {
		return param.Tokens{}, fmt.Errorf("couldn't create token: %w", jErr)
	}

	refreshToken, jErr := s.auth.CreateRefreshToken(user)
	if jErr != nil {
		return param.Tokens{}, fmt.Errorf("couldn't create token: %w", jErr)
	}

	return param.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
