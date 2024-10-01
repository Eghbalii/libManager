package main

import (
	"time"

	"github.com/eghbalii/libManager/httpserver"
	"github.com/eghbalii/libManager/repository/mongo"
	"github.com/eghbalii/libManager/repository/mongo/mongouser"
	"github.com/eghbalii/libManager/service/authservice"
	"github.com/eghbalii/libManager/service/userservice"
	"github.com/eghbalii/libManager/validator/uservalidator"
)

func main() {
	// http.ListenAndServe(":8090", nil)

	authSvc, userSvc, userValidator := setupService()
	server := httpserver.New(authSvc, userSvc, userValidator)

	server.Serve(8080)
}

func setupService() (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(authservice.Config{
		SignKey:              "jwt_secret",
		AccessExpirationTime: time.Hour * 24,
		AccessSubject:        "access",
	})
	mongoClient := mongo.NewClient(27018)
	mongoRepo := mongouser.New(mongoClient)
	userSvc := userservice.New(mongoRepo, authSvc)

	userValidator := uservalidator.New(mongoRepo)

	return authSvc, userSvc, userValidator
}
