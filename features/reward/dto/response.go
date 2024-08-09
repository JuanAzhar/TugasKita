package dto

type RewardResponse struct {
	Id    string
	Name  string
	Stock int
	Price int
	Image string
}

type RewardRequestResponse struct {
	Id         string
	RewardId   string
	RewardName string
	UserId     string
	UserNaame  string
	Status     string
}
