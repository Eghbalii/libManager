package bookhandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h Handler) ListBooks(e echo.Context) error {
	log.Println("<<<<<ListBook>>>>")
	books, err := h.bookSvc.ListBooks()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, books)
}

func (h Handler) GetBook(e echo.Context) error {
	log.Println("<<<<<GetBook>>>>")
	bookID := e.QueryParam("bookID")
	if bookID == "" {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "bookID is required"})
	}

	bookObjectID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "invalid bookID"})
	}
	book, err := h.bookSvc.GetBookByID(bookObjectID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, book)
}
