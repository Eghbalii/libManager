package main

import (
	"time"

	"github.com/eghbalii/libManager/delivery/httpserver"
	"github.com/eghbalii/libManager/repository/mongo"
	"github.com/eghbalii/libManager/repository/mongo/mongobook"
	"github.com/eghbalii/libManager/repository/mongo/mongosearch"
	"github.com/eghbalii/libManager/repository/mongo/mongouser"
	"github.com/eghbalii/libManager/service/authorizationservice"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/eghbalii/libManager/service/bookservice"
	"github.com/eghbalii/libManager/service/searchservice"
	"github.com/eghbalii/libManager/service/userservice"
	"github.com/eghbalii/libManager/validator/bookvalidator"
	"github.com/eghbalii/libManager/validator/uservalidator"
)

func main() {
	// http.ListenAndServe(":8090", nil)

	authSvc, authorizationSvc, userSvc, userValidator, bookSvc, bookValidator, searchSvc := setupService()
	server := httpserver.New(authSvc, authorizationSvc, userSvc, userValidator, bookSvc, bookValidator, searchSvc)

	server.Serve(8080)
}

func setupService() (authservice.Service, authorizationservice.Service, userservice.Service, uservalidator.Validator,
	bookservice.Service, bookvalidator.Validator, searchservice.Service) {
	mongoClient := mongo.NewClient(27018)
	authSvc := authservice.New(authservice.Config{
		SignKey:              "jwt_secret",
		AccessExpirationTime: time.Hour * 24,
		AccessSubject:        "access",
	})

	// user
	userMongoRepo := mongouser.New(mongoClient)
	userSvc := userservice.New(userMongoRepo, authSvc)
	userValidator := uservalidator.New(userMongoRepo)

	// book
	bookMongoRepo := mongobook.New(mongoClient)
	bookSvc := bookservice.New(bookMongoRepo)
	bookValidator := bookvalidator.New(bookMongoRepo)
	authorizationservice := authorizationservice.New(bookMongoRepo)

	// search
	searchMongoRepo := mongosearch.New(mongoClient)
	searchservice := searchservice.New(searchMongoRepo)

	return authSvc, authorizationservice, userSvc, userValidator, bookSvc, bookValidator, searchservice
}
