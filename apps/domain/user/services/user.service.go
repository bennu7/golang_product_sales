package services

import (
	"context"
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/bennu7/golang_product_sales/apps/domain/user/params"
	"github.com/bennu7/golang_product_sales/apps/domain/user/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{repo: repo}
}

func (u *UserService) ProfileMe(ctx context.Context, uid uuid.UUID) (*params.UserResponse, *responses.CustomError) {
	modelUser := u.repo.Profile(ctx, uid)
	responseUser := &params.UserResponse{}
	responseUser.ParseToModelResponse(modelUser)
	//finalData := modelUser
	if modelUser == nil {
		errResponse := responses.ErrResNotFound("user not found")
		return nil, errResponse
	}
	return responseUser, nil
}

func (u *UserService) DeleteMe(ctx context.Context, uid uuid.UUID) *responses.CustomError {
	err := u.repo.Delete(ctx, uid)
	if err != nil {
		errResponse := responses.ErrBadrequest("failed to delete user", err.Error())
		return errResponse
	}

	return nil
}
