package bookhandler

import (
	"context"
	"net/http"

	"github.com/eghbalii/libManager/contract/goproto/book"
	"github.com/labstack/echo/v4"
)

func (h Handler) BorrowBook(c echo.Context) error {
	req := new(book.BorrowBookRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	_, err := h.grpcClient.BorrowBook(context.Background(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "book borrowed successfully"})
}

func (h Handler) ReturnBook(c echo.Context) error {
	req := new(book.ReturnBookRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	_, err := h.grpcClient.ReturnBook(context.Background(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "book returned successfully"})
}

func (h Handler) ReserveBook(c echo.Context) error {
	req := new(book.ReserveBookRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	_, err := h.grpcClient.ReserveBook(context.Background(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "book reserved successfully"})
}
