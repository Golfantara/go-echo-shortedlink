package donate

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID uint `gorm:"primaryKey;type:int"`
	UserID    string `json:"user_id"`
	OrderID   string `gorm:"type:varchar(255)"`
	Status    string `gorm:"type:varchar(20);default:'pending'"`
	Amount int64
	Description string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Status struct {
	Transaction string
	Donate       string
}