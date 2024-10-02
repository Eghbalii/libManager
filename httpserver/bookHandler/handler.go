package bookhandler

import (
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/eghbalii/libManager/service/bookservice"
	"github.com/eghbalii/libManager/validator/bookvalidator"
)

type Handler struct {
	authSvc       authservice.Service
	bookSvc       bookservice.Service
	bookValidator bookvalidator.Validator
}

func New(authSvc authservice.Service, bookSvc bookservice.Service, bookValidator bookvalidator.Validator) Handler {
	return Handler{authSvc: authSvc, bookSvc: bookSvc, bookValidator: bookValidator}
}
