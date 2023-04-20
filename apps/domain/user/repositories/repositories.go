package repositories

import (
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

func (u *UserRepo) Profile(uid uuid.UUID) *models.User {
	userModel := &models.User{}
	tx := u.db.Debug().First(userModel, uid)
	if tx.RowsAffected < 1 {
		return nil
	}
	return userModel
}
