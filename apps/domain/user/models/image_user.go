package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ImageUser struct {
	Id        uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserId    uuid.UUID      `gorm:"type:uuid;unique" json:"user_id"`
	ImageUrl  string         `gorm:"not null" json:"image_url"`
	ImageKey  string         `gorm:"not null" json:"image_key"`
	CreatedAt *time.Time     `gorm:"autoCreateTime" json:"created_t"`
	UpdatedAt *time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"-"`
}
