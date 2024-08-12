package entity

import (
	"time"

	"github.com/google/uuid"
)

type PenaltyCore struct {
	Id          uuid.UUID
	UserId      string
	Point       int
	Description string
	Date        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
