package userservice

import (
	"context"

	"github.com/eghbalii/libManager/param"
)

func (s Service) Profile(ctx context.Context, req param.ProfileRequest) (param.ProfileResponse, error) {
	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return param.ProfileResponse{}, err
	}

	return param.ProfileResponse{Name: user.Name}, nil
}
