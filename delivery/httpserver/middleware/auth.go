package middleware

import (
	"github.com/javadifa/socialmedia/service/authservice"
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

//closure or higher order function --> it returns a function

func Auth(service authservice.Service, config authservice.Config) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey:    "user",
		SigningKey:    []byte(config.SignKey),
		SigningMethod: "HS256",
		KeyFunc:       nil,
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParsToken(auth)
			if err != nil {
				return nil, err
			}
			return claims, nil
		},
	})
}
