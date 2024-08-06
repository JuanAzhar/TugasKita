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

type RewardReqUpdateRequest struct {
	RewardId string `json:"reward_id"`
	UserId   string `json:"user_id"`
	Status   string `json:"status"`
}
