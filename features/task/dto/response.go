package dto

type TaskResponse struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Point      int    `json:"point"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	Start_date string `json:"startDate"`
	End_date   string `json:"endDate"`
}

type TaskResponseDetail struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Point       int    `json:"point"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	Start_date  string `json:"startDate"`
	End_date    string `json:"endDate"`
}

type UserTaskUploadResponse struct {
	Id          string
	TaskId      string
	UserId      string
	Image       string
	Description string
	Status      string
}
