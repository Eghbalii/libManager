package searchhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// SearchBookByTitle Search book by title
// @Summary Search book by title
// @Description Search book by title
// @Tags search
// @Accept json
// @Produce json
// @Param title query string true "title"
// @Success 200 {array} entity.Book
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /search/title [get]
func (h Handler) SearchBookByTitle(c echo.Context) error {
	title := c.QueryParam("title")
	books, err := h.serachSvc.SearchBookByTitle(title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, books)
}

// SearchBookByAuthor Search book by author
// @Summary Search book by author
// @Description Search book by author
// @Tags search
// @Accept json
// @Produce json
// @Param author query string true "author"
// @Success 200 {array} entity.Book
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /search/author [get]
func (h Handler) SearchBookByAuthor(c echo.Context) error {
	author := c.QueryParam("author")
	books, err := h.serachSvc.SearchBookByAuthor(author)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, books)
}
