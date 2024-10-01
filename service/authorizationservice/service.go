package authorizationservice

import (
	"github.com/eghbalii/libManager/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	GetUserPermissionTitles(userID primitive.ObjectID, role entity.Role) ([]entity.PermissionTitle, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) CheckAccess(userID primitive.ObjectID, role entity.Role, permissions ...entity.PermissionTitle) (bool, error) {
	permissionTitles, err := s.repo.GetUserPermissionTitles(userID, role)
	if err != nil {
		return false, err
	}

	for _, pt := range permissionTitles {
		for _, p := range permissions {
			if pt == p {
				return true, nil
			}
		}
	}

	return false, nil
}
