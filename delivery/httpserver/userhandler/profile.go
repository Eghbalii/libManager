package userHandler

import (
	"net/http"

	"github.com/eghbalii/libManager/param"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/labstack/echo/v4"
)

// UserProfile Get user profile
//
//	@Summary		Get user profile
//	@Description	Get user profile
//	@Tags			User
//	@Produce		json
//	@Success		200	{object}	param.ProfileResponse
//	@Router			/users/profile [get]
//
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h Handler) userProfile(c echo.Context) error {
	claims := c.Get("claims").(*authservice.Claims)
	resp, err := h.userSvc.Profile(c.Request().Context(),
		param.ProfileRequest{UserID: claims.UserID})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}
