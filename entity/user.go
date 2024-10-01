package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	PhoneNumber string             `bson:"phoneNumber"`
	Name        string             `bson:"name"`
	Password    string             `bson:"password"`
	Role        Role               `bson:"role"`
}
