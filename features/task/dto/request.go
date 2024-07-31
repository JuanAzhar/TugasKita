package dto

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Point       int    `json:"point"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"end_date"`
}

type UserTaskUploadRequest struct {
	TaskId      string `json:"task_id"`
	UserId      string `json:"user_id"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
