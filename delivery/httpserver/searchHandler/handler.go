package searchhandler

import (
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/eghbalii/libManager/service/searchservice"
)

type Handler struct {
	authSvc   authservice.Service
	serachSvc searchservice.Service
}

func New(authSvc authservice.Service, searchSvc searchservice.Service) Handler {
	return Handler{authSvc: authSvc, serachSvc: searchSvc}
}
