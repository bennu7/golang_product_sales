package repositories

import (
	"context"
	"fmt"
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepositoryInterface {
	return &UserRepo{db: db}
}

func (u *UserRepo) Profile(ctx context.Context, uid uuid.UUID) *models.User {
	userModel := &models.User{}
	tx := u.db.Debug().First(userModel, uid)
	if tx.RowsAffected < 1 {
		return nil
	}
	return userModel
}

func (u *UserRepo) Delete(ctx context.Context, uid uuid.UUID) error {
	userModel := &models.User{}
	tx := u.db.Debug().First(userModel, uid)
	if tx.RowsAffected < 1 {
		return fmt.Errorf("user not found")
	}

	tx = u.db.Debug().Delete(userModel, uid)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
