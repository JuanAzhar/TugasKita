package dto

type UserResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Point      string `json:"point"`
	TotalPoint string `json:"total_point"`
}

type UserRankResponse struct {
	Name  string `json:"name"`
	Point string `json:"point"`
}
