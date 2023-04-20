package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FullName  string     `gorm:"not null"`
	Email     string     `gorm:"not null;index:idx_email,unique"`
	Gender    string     `gorm:"not null"`
	Password  string     `gorm:"not null"`
	ImageUser ImageUser  `gorm:"foreignKey:UserId;associationForeignKey:Id;references:Id"`
	RoleId    uuid.UUID  `gorm:"null"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
