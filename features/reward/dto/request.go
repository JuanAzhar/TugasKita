package dto

type RewardRequest struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

type RewardReqRequest struct {
	RewardId string `json:"reward_id"`
}