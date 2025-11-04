package authservice

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/javadifa/socialmedia/entity"
)

//TODO: env for auth configuration(secret-key!!)

type Config struct {
	SignKey               string
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}

// in service small, in config with big letter
// fields we need
type Service struct {
	// no pointer cus we want a copy
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}

//todo: revoke tokens based on time

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.CreateToken(user.ID, s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.CreateToken(user.ID, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) ParsToken(bearerToken string) (*Claims, error) {
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, pErr := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})

	if pErr != nil {

		return nil, pErr
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		//2nd part is sus
		return nil, errors.New("invalid token")
	}

}

func (s Service) CreateToken(userID uint, subject string, expireDuration time.Duration) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, sErr := accessToken.SignedString([]byte(s.config.SignKey))

	if sErr != nil {
		return "", fmt.Errorf("unexpected error : %w", sErr)
	}
	return accessTokenString, nil
}
