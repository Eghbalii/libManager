package param

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProfileRequest struct {
	UserID primitive.ObjectID `json:"user_id"`
}

type ProfileResponse struct {
	Name string `json:"name"`
}
