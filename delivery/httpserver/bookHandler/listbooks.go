package bookhandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ListBooks List all books
// @Summary List all books
// @Description List all books
// @Tags book
// @Accept json
// @Produce json
// @Success 200 {array} entity.Book
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /books [get]
func (h Handler) ListBooks(e echo.Context) error {
	books, err := h.grpcClient.GetBooks(context.Background(), nil)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, books)
}
