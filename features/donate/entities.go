package donate

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID uint `gorm:"primaryKey;type:int"`
	// Order     model.Order `gorm:"foreignKey:ID"`
	OrderID   string `gorm:"type:varchar(255);autoIncrement"`
	Status    string `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Donated []Donated `gorm:"foreignKey:ID"`
}

type Status struct {
	Transaction string
	Donate       string
}

type Donated struct {
	ID    string
	Amount int64
}