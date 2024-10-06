package userHandler

import (
	"net/http"

	"github.com/eghbalii/libManager/param"
	"github.com/labstack/echo/v4"
)

// UserRegister Register user in system
//
//	@Summary		Register user in system
//	@Description	Get user data and register in system
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.RegisterRequest	true	"User data"
//	@Success		200		{object}	param.RegisterResponse
//	@Router			/users/register [post]
func (h Handler) userRegister(c echo.Context) error {
	var req param.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if filedErrors, err := h.userValidator.ValidateRegisterRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"errors": filedErrors})
	}
	resp, err := h.userSvc.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
