package searchhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) SearchBookByTitle(c echo.Context) error {
	title := c.QueryParam("title")
	books, err := h.serachSvc.SearchBookByTitle(title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, books)
}

func (h Handler) SearchBookByAuthor(c echo.Context) error {
	author := c.QueryParam("author")
	books, err := h.serachSvc.SearchBookByAuthor(author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, books)
}
