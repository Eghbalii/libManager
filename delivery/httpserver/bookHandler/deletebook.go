package bookhandler

import (
	"context"
	"net/http"

	"github.com/eghbalii/libManager/contract/goproto/book"
	"github.com/labstack/echo/v4"
)

// DeleteBook Delete book
// @Summary Delete book
// @Description Delete book by Admin
// @Tags book
// @Accept json
// @Produce json
// @Param bookID query string true "bookID"
// @Success 200 {object} string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /books/ [delete]
func (h Handler) DeleteBook(e echo.Context) error {
	bookID := e.QueryParam("bookID")
	if bookID == "" {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "bookID is required"})
	}

	_, err := h.grpcClient.DeleteBook(context.Background(), &book.DeleteBookRequest{BookID: bookID})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, echo.Map{"message": "book deleted successfully"})

}
