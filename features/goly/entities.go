package goly

import (
	"time"

	"gorm.io/gorm"
)

type Goly struct {
	ID        uint64 `gorm:"primaryKey"`
	Redirect  string `gorm:"not null"`
	Goly      string `gorm:"unique;not null"`
	Clicked   uint64 
	Random    bool   
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

}