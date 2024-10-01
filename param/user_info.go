package param

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserInfo struct {
	ID          primitive.ObjectID `json:"id"`
	PhoneNumber string             `json:"phone_number"`
	Name        string             `json:"name"`
}
