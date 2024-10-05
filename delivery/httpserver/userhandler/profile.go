package userHandler

import (
	"net/http"

	"github.com/eghbalii/libManager/param"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/labstack/echo/v4"
)

func (h Handler) userProfile(c echo.Context) error {
	claims := c.Get("claims").(*authservice.Claims)
	resp, err := h.userSvc.Profile(c.Request().Context(),
		param.ProfileRequest{UserID: claims.UserID})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}
