package dto

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Image    string `json:"image" form:"image"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
