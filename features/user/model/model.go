package model

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID         uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Name       string    `gorm:"varchar(50);not null" json:"username"`
	Email      string    `gorm:"varchar(50);not null" json:"email"`
	Password   string    `gorm:"varchar(50);not null" json:"password"`
	Role       string    `gorm:"Varchar(25);not null" json:"role"`
	Point      string    `gorm:"Varchar(100);not null" json:"point"`
	TotalPoint string    `gorm:"Varchar(100);not null" json:"total_point"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"update_at"`
}
