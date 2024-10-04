package bookhandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) ListBooks(e echo.Context) error {
	books, err := h.grpcClient.GetBooks(context.Background(), nil)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, books)
}
