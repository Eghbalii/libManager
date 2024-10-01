package authservice

import (
	"github.com/eghbalii/libManager/entity"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID primitive.ObjectID `json:"user_id"`
	Role   entity.Role        `json:"role"`
}

func (c Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}
