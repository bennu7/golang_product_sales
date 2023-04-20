package services

import (
	"context"
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/dto"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/repositories"
	"github.com/bennu7/golang_product_sales/pkg/encryption"
	"github.com/bennu7/golang_product_sales/pkg/token"
)

type authSvc struct {
	repo repositories.AuthRepo
}

func NewAuthSvc(repo repositories.AuthRepo) AuthServiceInterface {
	return &authSvc{repo: repo}
}

func (a *authSvc) RegisterSvc(ctx context.Context, dto *dto.AuthRegisterDTO) *responses.CustomError {
	if err := a.repo.CheckEmailRegister(ctx, dto.Email); err != nil {
		errResponse := responses.CustomError{
			StatusCode: 400,
			Message:    err.Error(),
		}
		return &errResponse
	}

	hashPass, _ := encryption.HashPassword(dto.Password)
	dto.Password = hashPass

	err := a.repo.Register(ctx, dto)
	if err != nil {
		errResponse := &responses.CustomError{
			StatusCode: 400,
			Message:    err.Error(),
		}
		return errResponse
	}

	return nil
}

func (a *authSvc) LoginSvc(ctx context.Context, dto *dto.AuthLoginDTO) (string, *responses.CustomError) {
	dataUser, err := a.repo.CheckEmailLogin(ctx, dto.Email)
	if err != nil {
		errResponse := &responses.CustomError{
			StatusCode: 404,
			Message:    err.Error(),
		}
		return "", errResponse
	}

	err = encryption.ValidatePassword(dataUser.Password, dto.Password)
	if err != nil {
		errResponse := &responses.CustomError{
			StatusCode: 404,
			Message:    "password not match",
		}
		return "", errResponse
	}

	payload := &token.Payload{
		UserId: dataUser.Id,
	}
	generateToken, err := token.GenerateToken(payload)
	if err != nil {
		errResponse := &responses.CustomError{
			StatusCode: 400,
			Message:    err.Error(),
		}
		return "", errResponse
	}

	return generateToken, nil
}
