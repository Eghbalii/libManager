package httpserver

import (
	"fmt"
	"log"

	"github.com/eghbalii/libManager/httpserver/userhandler"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/eghbalii/libManager/service/userservice"
	"github.com/eghbalii/libManager/validator/uservalidator"
	"github.com/labstack/echo/v4"
	md "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	userHandler userhandler.Handler
	Router      *echo.Echo
}

func New(authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator) Server {
	e := echo.New()
	userHandler := userhandler.New(authSvc, userSvc, userValidator)
	return Server{
		userHandler: userHandler,
		Router:      e,
	}
}

func (s Server) Serve(port int) {
	s.Router.Use(md.RequestID())
	s.Router.Use(md.Logger())
	s.Router.Use(md.Recover())

	s.userHandler.SetRoures(s.Router)

	// start server
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("echo server started at %s\n", address)
	if err := s.Router.Start(address); err != nil {
		log.Fatalf("failed to start router: %v", err)
	}
}
