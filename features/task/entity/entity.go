package entity

import (
	"time"

	"github.com/google/uuid"
)

type TaskCore struct {
	ID          uuid.UUID
	AdminId     string
	Title       string
	Description string
	Point       int
	Message     string
	Status      string
	Type        string
	Start_date  string
	End_date    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserTaskUploadCore struct {
	Id          uuid.UUID
	TaskId      string
	TaskName    string
	UserId      string
	UserName    string
	Image       string
	Description string
	Status      string
}

type UserTaskSubmissionCore struct {
	Id          uuid.UUID
	Title       string
	UserId      string
	Image       string
	Description string
	Point       int
	Status      string
}
