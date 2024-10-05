package bookhandler

import (
	"log"

	"github.com/eghbalii/libManager/contract/goproto/book"
	"github.com/eghbalii/libManager/service/authorizationservice"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/eghbalii/libManager/service/bookservice"
	"github.com/eghbalii/libManager/validator/bookvalidator"
	"google.golang.org/grpc"
)

type Handler struct {
	authSvc              authservice.Service
	authorizationservice authorizationservice.Service
	grpcClient           book.BookServiceClient
	bookValidator        bookvalidator.Validator
}

func New(authSvc authservice.Service, authorizationservice authorizationservice.Service, bookSvc bookservice.Service,
	bookValidator bookvalidator.Validator) Handler {
	conn, err := grpc.NewClient(":8086", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc client can not connect: %v", err)
	}
	grpcClient := book.NewBookServiceClient(conn)
	return Handler{authSvc: authSvc, authorizationservice: authorizationservice, grpcClient: grpcClient, bookValidator: bookValidator}
}
