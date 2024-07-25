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
	Start_date  time.Time
	End_date    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
