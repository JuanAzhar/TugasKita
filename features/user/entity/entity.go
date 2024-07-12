package entity

import "time"

type UserCore struct {
	ID         string    `json:"id"`
	Name       string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Point      string    `json:"point"`
	TotalPoint string    `gorm:"Varchar(100);not null;default:user" json:"total_point"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"update_at"`
}
