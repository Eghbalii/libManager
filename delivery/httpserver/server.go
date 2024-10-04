package httpserver

import (
	"fmt"
	"log"

	bookhandler "github.com/eghbalii/libManager/delivery/httpserver/bookHandler"
	"github.com/eghbalii/libManager/delivery/httpserver/userhandler"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/eghbalii/libManager/service/bookservice"
	"github.com/eghbalii/libManager/service/userservice"
	"github.com/eghbalii/libManager/validator/bookvalidator"
	"github.com/eghbalii/libManager/validator/uservalidator"
	"github.com/labstack/echo/v4"
	md "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	userHandler userhandler.Handler
	bookHandler bookhandler.Handler
	Router      *echo.Echo
}

func New(authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator,
	bookSvc bookservice.Service, bookValidator bookvalidator.Validator) Server {
	e := echo.New()
	userHandler := userhandler.New(authSvc, userSvc, userValidator)
	bookHandler := bookhandler.New(authSvc, bookSvc, bookValidator)
	return Server{
		userHandler: userHandler,
		bookHandler: bookHandler,
		Router:      e,
	}
}

func (s Server) Serve(port int) {
	s.Router.Use(md.RequestID())
	s.Router.Use(md.Logger())
	s.Router.Use(md.Recover())

	s.userHandler.SetRoutes(s.Router)
	s.bookHandler.SetRoutes(s.Router)

	// start server
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("echo server started at %s\n", address)
	if err := s.Router.Start(address); err != nil {
		log.Fatalf("failed to start router: %v", err)
	}
}
