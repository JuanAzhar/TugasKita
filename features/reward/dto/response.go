package dto

type RewardResponse struct {
	Id    string
	Name  string
	Stock int
	Price int
	Image string
}

type RewardRequestResponse struct {
	Id       string
	RewardId string
	UserId   string
	Status   string
}
