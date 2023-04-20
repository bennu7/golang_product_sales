package services

import (
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
	"github.com/bennu7/golang_product_sales/apps/domain/user/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{repo: repo}
}

func (u *UserService) ProfileMe(uid uuid.UUID) *models.User {
	modelUser := u.repo.Profile(uid)
	if modelUser == nil {
		return nil
	}
	return modelUser
}
