package entity

import (
	"time"

	"github.com/google/uuid"
)

type RewardCore struct {
	ID        uuid.UUID
	Name      string
	Stock     int
	Price     int
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRewardRequestCore struct {
	Id         uuid.UUID
	RewardId   string
	RewardName string
	UserId     string
	UserName   string
	Status     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
