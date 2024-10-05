package userHandler

import (
	"net/http"

	"github.com/eghbalii/libManager/param"
	"github.com/labstack/echo/v4"
)

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
