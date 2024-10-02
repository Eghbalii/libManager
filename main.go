package main

import (
	"time"

	"github.com/eghbalii/libManager/httpserver"
	"github.com/eghbalii/libManager/repository/mongo"
	"github.com/eghbalii/libManager/repository/mongo/mongobook"
	"github.com/eghbalii/libManager/repository/mongo/mongouser"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/eghbalii/libManager/service/bookservice"
	"github.com/eghbalii/libManager/service/userservice"
	"github.com/eghbalii/libManager/validator/bookvalidator"
	"github.com/eghbalii/libManager/validator/uservalidator"
)

func main() {
	// http.ListenAndServe(":8090", nil)

	authSvc, userSvc, userValidator, bookSvc, bookValidator := setupService()
	server := httpserver.New(authSvc, userSvc, userValidator, bookSvc, bookValidator)

	server.Serve(8080)
}

func setupService() (authservice.Service, userservice.Service, uservalidator.Validator,
	bookservice.Service, bookvalidator.Validator) {
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

	return authSvc, userSvc, userValidator, bookSvc, bookValidator
}
