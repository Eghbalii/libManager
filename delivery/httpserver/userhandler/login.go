package userHandler

import (
	"net/http"

	"github.com/eghbalii/libManager/param"
	"github.com/labstack/echo/v4"
)

// UserLogin Login user in system
// @Summary		Login user in system
// @Description	Get user data and login in system
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body		param.LoginRequest	true	"User data"
// @Success		200		{object}	param.LoginResponse
// @Router			/users/login [post]
func (h Handler) userlLogin(c echo.Context) error {
	var req param.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if filedErrors, err := h.userValidator.ValidateLoginRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, filedErrors)
	}

	resp, err := h.userSvc.Login(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
