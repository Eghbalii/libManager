package userservice

import (
	"github.com/eghbalii/libManager/entity"
	"github.com/eghbalii/libManager/param"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Service) Register(req param.RegisterRequest) (param.RegisterResponse, error) {
	user := entity.User{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    getMD5Hash(req.Password),
		Role:        entity.UserRole,
	}

	// create new user in database
	user, err := s.repo.Register(user)
	if err != nil {
		return param.RegisterResponse{}, err
	}

	return param.RegisterResponse{
		User: param.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
	}, nil
}
