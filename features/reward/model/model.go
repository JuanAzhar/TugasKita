package model

import (
	"time"

	"github.com/google/uuid"
)

type Reward struct {
	ID        uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Name      string
	Stock     int
	Price     int
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRewardRequest struct {
	Id       uuid.UUID
	RewardId string 
	UserId   string
	Status   string `gorm:"type:varchar(20);default:'Review'" json:"status"`
}
