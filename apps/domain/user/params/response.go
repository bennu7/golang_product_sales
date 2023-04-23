package params

import (
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
	"github.com/google/uuid"
	"time"
)

type UserResponse struct {
	Id        uuid.UUID        `json:"id"`
	FullName  string           `json:"full_name"`
	Email     string           `json:"email"`
	Gender    string           `json:"gender"`
	ImageUser models.ImageUser `json:"image_user"`
	RoleId    uuid.UUID        `json:"role_id"`
	CreatedAt *time.Time       `json:"created_at"`
	UpdatedAt *time.Time       `json:"updated_at"`
}

func (u *UserResponse) ParseToModelResponse(user *models.User) {
	u.Id = user.Id
	u.FullName = user.FullName
	u.Email = user.Email
	u.Gender = user.Gender
	u.ImageUser = user.ImageUser
	u.RoleId = user.RoleId
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
}

type ImageUserResponse struct{}
