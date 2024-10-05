package bookhandler

import (
	"github.com/eghbalii/libManager/delivery/httpserver/middleware"
	"github.com/eghbalii/libManager/entity"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	e.POST("/books", h.AddBook, middleware.Auth(h.authSvc))
	e.GET("/books", h.ListBooks, middleware.Auth(h.authSvc))
	e.PUT("/books/", h.UpdateBook, middleware.Auth(h.authSvc),
		middleware.AccessCheck(h.authorizationservice, entity.BookUpdatePermission))
	e.DELETE("/books/", h.DeleteBook, middleware.Auth(h.authSvc),
		middleware.AccessCheck(h.authorizationservice, entity.BookDeletePermission))
}
