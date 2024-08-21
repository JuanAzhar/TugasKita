package dto

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Image    string `json:"image" form:"image"`
	Address  string `json:"address"`
	School   string `json:"school"`
	Class    string `json:"class"`
	Religion string `json:"religion" form:"religion"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Point    string `json:"point" form:"point"`
}
