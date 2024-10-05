package mongobook

import (
	"context"
	"fmt"

	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *DB) GetUserPermissionTitles(userID primitive.ObjectID) ([]entity.PermissionTitle, error) {
	userCollection := d.conn.Database("libManager").Collection("users")
	filter := bson.M{"_id": userID}
	var user entity.User
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	if user.Role != entity.AdminRole {
		return nil, fmt.Errorf("user role is not Admin")
	}

	return []entity.PermissionTitle{
		entity.BookUpdatePermission,
		entity.BookDeletePermission,
	}, nil
}
