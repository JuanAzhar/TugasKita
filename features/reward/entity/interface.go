package entity

type RewardDataInterface interface {
	CreateReward(input RewardCore) error
	FindAllReward() ([]RewardCore, error)
	FindById(rewardId string) (RewardCore, error)
	UpdateReward(rewardId string, data RewardCore) error
	DeleteReward(rewardId string) error

	UploadRewardRequest(input UserRewardRequestCore) error
	FindAllUploadReward()([]UserRewardRequestCore, error)
	FindUserRewardById(id string)(UserRewardRequestCore, error)
	FindAllRewardHistory(userId string)([]UserRewardRequestCore, error)
}

type RewardUseCaseInterface interface {
	CreateReward(input RewardCore) error
	FindAllReward() ([]RewardCore, error)
	FindById(rewardId string) (RewardCore, error)
	UpdateReward(rewardId string, data RewardCore) error
	DeleteReward(rewardId string) error

	UploadRewardRequest(input UserRewardRequestCore) error
	FindAllUploadReward()([]UserRewardRequestCore, error)
	FindUserRewardById(id string)(UserRewardRequestCore, error)
	FindAllRewardHistory(userId string)([]UserRewardRequestCore, error)
}
