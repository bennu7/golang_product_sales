package repositories

import (
	"context"
	"fmt"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/dto"
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{db: db}
}

func (a *authRepo) Register(ctx context.Context, registerUser *dto.AuthRegisterDTO) error {
	userModel := &models.User{}
	dtoDta := registerUser.ParseToModel()

	tx := a.db.Debug().Model(userModel).Create(dtoDta)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (a *authRepo) CheckEmailRegister(ctx context.Context, email string) error {
	dataUser := &models.User{}
	tx := a.db.Debug().Where("email = ?", email).First(dataUser)
	if tx.RowsAffected > 0 {
		return fmt.Errorf("email already exists")
	}
	return nil
}

func (a *authRepo) CheckEmailLogin(ctx context.Context, email string) (*models.User, error) {
	userModel := &models.User{}
	tx := a.db.Debug().Where("email = ?", email).First(userModel)
	if tx.RowsAffected < 1 {
		return userModel, fmt.Errorf("email not found")
	}
	return userModel, nil
}
