package bookhandler

import (
	"context"
	"net/http"

	"github.com/eghbalii/libManager/contract/goproto/book"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/labstack/echo/v4"
)

// BorrowBook Borrow a book
// @Summary Borrow a book
// @Description Borrow a book
// @Tags book
// @Accept json
// @Produce json
// @Param bookID body string true "bookID"
// @Success 200 {string} message
// Security JWT
// @Router /books/borrow [post]
func (h Handler) BorrowBook(c echo.Context) error {
	req := new(book.BorrowBookRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Retrieve the claims from the context
	claims := c.Get("claims").(*authservice.Claims)

	req.UserID = claims.UserID.Hex()

	isValidate, err := h.bookValidator.ValidateBorrowBookRequest(req.BookID, req.UserID)
	if err != nil || !isValidate {
		return c.JSON(http.StatusBadRequest, echo.Map{"borrow validation error": err.Error()})
	}

	_, err = h.grpcClient.BorrowBook(context.Background(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "book borrowed successfully"})
}

// ReturnBook Return a book
// @Summary Return a book
// @Description Return a book
// @Tags book
// @Accept json
// @Produce json
// @Param bookID body string true "bookID"
// @Success 200 {string} message
// Security JWT
// @Router /books/return [post]
func (h Handler) ReturnBook(c echo.Context) error {
	req := new(book.ReturnBookRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Retrieve the claims from the context
	claims := c.Get("claims").(*authservice.Claims)
	userID := claims.UserID.Hex()

	isValidate, err := h.bookValidator.ValidateReturnBookRequest(req.BookID, userID)
	if err != nil || !isValidate {
		return c.JSON(http.StatusBadRequest, echo.Map{"return validation error": err.Error()})
	}

	_, err = h.grpcClient.ReturnBook(context.Background(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "book returned successfully"})
}

// ReserveBook Reserve a book
// @Summary Reserve a book
// @Description Reserve a book
// @Tags book
// @Accept json
// @Produce json
// @Param bookID body string true "bookID"
// @Success 200 {string} message
// Security JWT
// @Router /books/reserve [post]
func (h Handler) ReserveBook(c echo.Context) error {
	req := new(book.ReserveBookRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Retrieve the claims from the context
	claims := c.Get("claims").(*authservice.Claims)

	// Extract the userID from the claims
	userID := claims.UserID.Hex()
	req.UserID = userID

	isValid, err := h.bookValidator.ValidateReserveBookRequest(req.BookID, req.UserID)
	if err != nil || !isValid {
		return c.JSON(http.StatusBadRequest, echo.Map{"reserve validation error": err.Error()})
	}

	_, err = h.grpcClient.ReserveBook(context.Background(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "book reserved successfully"})
}
