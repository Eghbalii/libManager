package bookhandler

import (
	"net/http"

	"github.com/eghbalii/libManager/param"
	"github.com/labstack/echo/v4"
)

func (h Handler) UpdateBook(e echo.Context) error {
	var req param.UpdateBookRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	if fieldErrors, err := h.bookValidator.ValidateUpdateBookRequest(req); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"errors": fieldErrors})
	}
	book, err := h.bookSvc.UpdateBook(req)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, book)
}
