package bookhandler

import (
	"github.com/eghbalii/libManager/httpserver/middleware"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	e.POST("/books", h.AddBook, middleware.Auth(h.authSvc))
	e.GET("/books", h.ListBooks, middleware.Auth(h.authSvc))
	e.GET("/books/", h.GetBook, middleware.Auth(h.authSvc))
	e.PUT("/books/", h.UpdateBook, middleware.Auth(h.authSvc))
	e.DELETE("/books/", h.DeleteBook, middleware.Auth(h.authSvc))
}
