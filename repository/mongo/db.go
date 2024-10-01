package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(port int) *mongo.Client {
	url := fmt.Sprintf("mongodb://localhost:%d", port)
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	// connect to mongo
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
