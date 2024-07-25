package dto

import "time"

type TaskResponse struct {
	Title       string    `json:"title"`
	Point       int       `json:"point"`
	Start_date  time.Time `json:"startDate"`
	End_date    time.Time `json:"endDate"`
}

type TaskResponseDetail struct{
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Point       int       `json:"point"`
	Message     string    `json:"message"`
	Status      string    `json:"status"`
	Start_date  time.Time `json:"startDate"`
	End_date    time.Time `json:"endDate"`
}
