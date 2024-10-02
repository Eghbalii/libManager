package mongobook

import "go.mongodb.org/mongo-driver/mongo"

type DB struct {
	conn *mongo.Client
}

func New(conn *mongo.Client) *DB {
	return &DB{conn: conn}
}
