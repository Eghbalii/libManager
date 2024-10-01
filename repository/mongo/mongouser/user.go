package mongouser

import (
	"context"
	"fmt"

	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	// search in database for phoneNumber
	userCollection := d.conn.Database("libManager").Collection("users")

	filter := bson.M{"phoneNumber": phoneNumber}
	result := userCollection.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		return false, err
	}
	return false, fmt.Errorf("phone number is already exists")
}

func (d *DB) Register(u entity.User) (entity.User, error) {
	userCollection := d.conn.Database("libManager").Collection("users")
	_, err := userCollection.InsertOne(context.Background(), u)
	if err != nil {
		return entity.User{}, err
	}
	return u, nil
}

func (d *DB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	userCollection := d.conn.Database("libManager").Collection("users")
	filter := bson.M{"phoneNumber": phoneNumber}
	result := userCollection.FindOne(context.Background(), filter)
	if err := result.Err(); err != nil {
		return entity.User{}, err
	}
	var user entity.User
	err := result.Decode(&user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (d *DB) GetUserByID(ctx context.Context, userID primitive.ObjectID) (entity.User, error) {
	userCollection := d.conn.Database("libManager").Collection("users")
	filter := bson.M{"_id": userID}
	result := userCollection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return entity.User{}, err
	}
	var user entity.User
	err := result.Decode(&user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
