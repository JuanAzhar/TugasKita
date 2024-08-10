package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	AdminId     string
	Title       string `gorm:"not null" json:"title"`
	Description string
	Point       int
	Message     string
	Status      string `gorm:"type:varchar(20);default:'Active'" json:"status"`
	Type        string `gorm:"default:'Task'" json:"type"`
	Start_date  string
	End_date    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserTaskUpload struct {
	Id          uuid.UUID
	TaskId      string
	UserId      string
	Image       string
	Description string
	Status      string `gorm:"type:varchar(20);default:'Perlu Review'" json:"status"`
	Message     string
}

type UserTaskSubmission struct {
	Id          uuid.UUID
	Title       string
	UserId      string
	Image       string
	Type        string `gorm:"default:'Submission'" json:"type"`
	Description string
	Point       int
	Status      string `gorm:"type:varchar(20);default:'Perlu Review'" json:"status"`
	Message     string
}
