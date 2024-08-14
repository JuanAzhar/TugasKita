package model

import "time"

type UserPoint struct {
	Id        string `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	UserId    string
	Type      string
	TaskName  string
	Point     int
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
