package httpserver

import (
	"fmt"
	"log"

	bookhandler "github.com/eghbalii/libManager/delivery/httpserver/bookHandler"
	searchhandler "github.com/eghbalii/libManager/delivery/httpserver/searchHandler"
	"github.com/eghbalii/libManager/delivery/httpserver/userHandler"
	"github.com/eghbalii/libManager/service/authorizationservice"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/eghbalii/libManager/service/bookservice"
	"github.com/eghbalii/libManager/service/searchservice"
	"github.com/eghbalii/libManager/service/userservice"
	"github.com/eghbalii/libManager/validator/bookvalidator"
	"github.com/eghbalii/libManager/validator/uservalidator"
	"github.com/labstack/echo/v4"
	md "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	userHandler   userHandler.Handler
	bookHandler   bookhandler.Handler
	searchHandler searchhandler.Handler
	Router        *echo.Echo
}

func New(authSvc authservice.Service, authorizationSvc authorizationservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator,
	bookSvc bookservice.Service, bookValidator bookvalidator.Validator, searchSvc searchservice.Service) Server {
	e := echo.New()
	userHandler := userHandler.New(authSvc, userSvc, userValidator)
	bookHandler := bookhandler.New(authSvc, authorizationSvc, bookSvc, bookValidator)
	searchHandler := searchhandler.New(authSvc, searchSvc)
	return Server{
		userHandler:   userHandler,
		bookHandler:   bookHandler,
		searchHandler: searchHandler,
		Router:        e,
	}
}

func (s Server) Serve(port int) {
	s.Router.Use(md.RequestID())
	s.Router.Use(md.Logger())
	s.Router.Use(md.Recover())

	s.userHandler.SetRoutes(s.Router)
	s.bookHandler.SetRoutes(s.Router)
	s.searchHandler.SetRoutes(s.Router)

	// start server
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("echo server started at %s\n", address)
	if err := s.Router.Start(address); err != nil {
		log.Fatalf("failed to start router: %v", err)
	}
}
