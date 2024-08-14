package entity

import "time"

type UserPoint struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Type      string    `json:"type"`
	TaskName  string    `json:"task_name"`
	Point     int       `json:"point"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
