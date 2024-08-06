package entity

import (
	"time"
	"tugaskita/features/reward/model"
)

func RewardCoreToRewardModel(data RewardCore) model.Reward {
	return model.Reward{
		ID:    data.ID,
		Name:  data.Name,
		Stock: data.Stock,
		Price: data.Price,
		Image: data.Image,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func RewardModelToRewardCore(data model.Reward) RewardCore {
	return RewardCore{
		ID:    data.ID,
		Name:  data.Name,
		Stock: data.Stock,
		Price: data.Price,
		Image: data.Image,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func ListRewardModelToRewardCore(data []model.Reward) []RewardCore {
	dataReward := []RewardCore{}
	for _, v := range data {
		result := RewardModelToRewardCore(v)
		dataReward = append(dataReward, result)
	}
	return dataReward
}

func RewardUserModelToRewardUserCore(data model.UserRewardRequest) UserRewardRequestCore{
	return UserRewardRequestCore{
		Id: data.Id,
		RewardId: data.RewardId,
		UserId: data.UserId,
		Status: data.Status,
	}
}

func RewardUserCoreToRewardUserModel(data UserRewardRequestCore) model.UserRewardRequest{
	return model.UserRewardRequest{
		Id: data.Id,
		RewardId: data.RewardId,
		UserId: data.UserId,
		Status: data.Status,
	}
} 

func ListRewardUserModelToListRewardUserCore(data []model.UserRewardRequest) []UserRewardRequestCore{
	dataReward := []UserRewardRequestCore{}
	for _, v := range data {
		result := RewardUserModelToRewardUserCore(v)
		dataReward = append(dataReward, result)
	}
	return dataReward
}