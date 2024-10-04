package middleware

import (
	"net/http"

	"github.com/eghbalii/libManager/entity"
	"github.com/eghbalii/libManager/service/authorizationservice"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/labstack/echo/v4"
)

func AccessCheck(service authorizationservice.Service,
	permissions ...entity.PermissionTitle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("claims").(*authservice.Claims)
			isAllowed, err := service.CheckAccess(claims.UserID, claims.Role, permissions...)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			if !isAllowed {
				return c.JSON(http.StatusForbidden, "You don't have permission to access this resource")
			}
			return next(c)
		}
	}
}
