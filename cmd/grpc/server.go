package main

import (
	"github.com/eghbalii/libManager/delivery/grpcserver"
	"github.com/eghbalii/libManager/repository/mongo"
	"github.com/eghbalii/libManager/repository/mongo/mongobook"
	"github.com/eghbalii/libManager/service/bookservice"
)

func main() {
	mongoClient := mongo.NewClient(27018)

	bookMongoRepo := mongobook.New(mongoClient)
	bookSvc := bookservice.New(bookMongoRepo)

	server := grpcserver.New(bookSvc)
	server.Start()

}
