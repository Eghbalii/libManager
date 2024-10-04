package bookhandler

import (
	"context"
	"net/http"

	"github.com/eghbalii/libManager/contract/goproto/book"
	"github.com/labstack/echo/v4"
)

func (h Handler) AddBook(e echo.Context) error {
	req := new(book.Book)
	if err := e.Bind(req); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	if fieldErrors, err := h.bookValidator.ValidateAddBookRequest(*req); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"errors": fieldErrors})
	}
	book, err := h.grpcClient.CreateBook(context.Background(), req)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusCreated, book)
}
