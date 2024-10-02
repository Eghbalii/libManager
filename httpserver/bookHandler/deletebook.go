package bookhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h Handler) DeleteBook(e echo.Context) error {
	bookID := e.QueryParam("bookID")
	if bookID == "" {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "bookID is required"})
	}
	bookObjectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "invalid bookID"})
	}
	err = h.bookSvc.DeleteBook(bookObjectID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, echo.Map{"message": "book deleted successfully"})

}
