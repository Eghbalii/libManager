package searchhandler

import (
	"github.com/eghbalii/libManager/delivery/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	e.GET("/search/book/title", h.SearchBookByTitle, middleware.Auth(h.authSvc))
	e.GET("/search/book/author", h.SearchBookByAuthor, middleware.Auth(h.authSvc))
}
