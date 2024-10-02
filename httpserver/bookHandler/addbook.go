package bookhandler

import (
	"net/http"

	"github.com/eghbalii/libManager/param"
	"github.com/labstack/echo/v4"
)

func (h Handler) AddBook(e echo.Context) error {
	var req param.AddBookRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	if fieldErrors, err := h.bookValidator.ValidateAddBookRequest(req); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"errors": fieldErrors})
	}
	book, err := h.bookSvc.AddBook(req)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusCreated, book)
}
