package bookhandler

import (
	"context"
	"net/http"

	"github.com/eghbalii/libManager/contract/goproto/book"
	"github.com/labstack/echo/v4"
)

// UpdateBook Update book
// @Summary Update book
// @Description Update book by Admin
// @Tags book
// @Accept json
// @Produce json
// @Param book body entity.Book true "book"
// @Success 200 {object} entity.Book
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /books/ [put]
func (h Handler) UpdateBook(e echo.Context) error {
	req := new(book.Book)
	if err := e.Bind(req); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	if fieldErrors, err := h.bookValidator.ValidateUpdateBookRequest(*req); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"errors": fieldErrors})
	}
	book, err := h.grpcClient.UpdateBook(context.Background(), req)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, book)
}
