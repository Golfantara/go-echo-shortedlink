package auth

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID          int    `gorm:"primaryKey";"type:int(11); "`
	Fullname    string `gorm:"type:varchar(255);not null"`
	PhoneNumber string `gorm:"type:varchar(13);uniqueIndex"`
	Email       string `gorm:"type:varchar(255);not null"`
	Password    string `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}