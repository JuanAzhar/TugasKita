package dto

import "time"

type TaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Point       int       `json:"point"`
	Message     string    `json:"message"`
	Status      string    `json:"status"`
	Start_date  time.Time `json:"startDate"`
	End_date    time.Time `json:"endDate"`
}
