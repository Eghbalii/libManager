package middleware

import (
	"github.com/eghbalii/libManager/service/authservice"
	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Auth(service authservice.Service) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey:    "claims",
		SigningKey:    []byte("jwt_secret"),
		SigningMethod: "HS256",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			// Check for both Authorization header and Bearer token in the request body
			token := auth
			if token == "" && c.Request().Header.Get("Authorization") != "" {
				token = c.Request().Header.Get("Authorization")
			}

			claims, err := service.ParseToken(token)
			if err != nil {
				return nil, err
			}
			return claims, nil
		},
	})
}
