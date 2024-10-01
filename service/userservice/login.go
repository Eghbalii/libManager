package userservice

import (
	"fmt"

	"github.com/eghbalii/libManager/param"
)

func (s Service) Login(req param.LoginRequest) (param.LoginResponse, error) {
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, err
	}

	if user.Password != getMD5Hash(req.Password) {
		return param.LoginResponse{}, fmt.Errorf("username or password is incorrect")
	}

	accessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return param.LoginResponse{}, err
	}

	return param.LoginResponse{
		User: param.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		AccessToken: accessToken,
	}, nil
}
