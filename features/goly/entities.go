package goly

import (
	"time"

	"gorm.io/gorm"
)

type Goly struct {
	ID        uint64 `gorm:"primaryKey"`
	UserID    string `json:"user_id"`
	Redirect  string `gorm:"not null"`
	Custom    string `gorm:"unique;not null"`
	Clicked   uint64 
	Random    bool   
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IPAdresses []IPAdresses `gorm:"foreignKey:GolyID"`
}

type IPAdresses struct {
	ID		uint64
	GolyID	uint64
	Address	string
	CreatedAt time.Time
}