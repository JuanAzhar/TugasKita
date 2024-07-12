package handler

type UserResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Point      string `json:"point"`
	TotalPoint string `json:"total_point"`
}
