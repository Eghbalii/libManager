package userhandler

import (
	"github.com/eghbalii/libManager/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.POST("/login", h.userlLogin)
	userGroup.POST("/register", h.userRegister)
	userGroup.GET("/profile", h.userProfile,
		middleware.Auth(h.authSvc))
}
