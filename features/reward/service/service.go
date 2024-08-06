package service

import (
	"errors"
	"strconv"
	"tugaskita/features/reward/entity"
	user "tugaskita/features/user/entity"
)

type RewardService struct {
	RewardRepo entity.RewardDataInterface
	UserRepo user.UserDataInterface
}

func NewRewardService(rewardRepo entity.RewardDataInterface, userRepo user.UserDataInterface) entity.RewardUseCaseInterface {
	return &RewardService{
		RewardRepo: rewardRepo,
		UserRepo: userRepo,
	}
}

// CreateReward implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) CreateReward(input entity.RewardCore) error {

	if input.Name == "" || input.Image == "" {
		return errors.New("name and image can't be empty")
	}

	if input.Price < 0 || input.Stock < 0 {
		return errors.New("price and stock can't less then 0")
	}

	err := rewardUC.RewardRepo.CreateReward(input)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) DeleteReward(rewardId string) error {
	if rewardId == "" {
		return errors.New("insert reward id")
	}

	_, err := rewardUC.RewardRepo.FindById(rewardId)
	if err != nil {
		return errors.New("reward not found")
	}

	errDelete := rewardUC.RewardRepo.DeleteReward(rewardId)
	if errDelete != nil {
		return errors.New("can't delete reward")
	}

	return nil
}

// FindAllReward implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) FindAllReward() ([]entity.RewardCore, error) {
	data, err := rewardUC.RewardRepo.FindAllReward()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return data, nil
}

// FindById implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) FindById(rewardId string) (entity.RewardCore, error) {
	if rewardId == "" {
		return entity.RewardCore{}, errors.New("reward ID is required")
	}

	task, err := rewardUC.RewardRepo.FindById(rewardId)
	if err != nil {
		return entity.RewardCore{}, err
	}

	return task, nil
}

// UpdateReward implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) UpdateReward(rewardId string, data entity.RewardCore) error {
	if data.Name == "" || data.Image == "" {
		return errors.New("name and image can't be empty")
	}

	if data.Price < 0 || data.Stock < 0 {
		return errors.New("price and stock can't less then 0")
	}

	err := rewardUC.RewardRepo.UpdateReward(rewardId, data)
	if err != nil {
		return err
	}

	return nil
}

// FindAllRewardHistory implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) FindAllRewardHistory(userId string) ([]entity.UserRewardRequestCore, error) {
	data, err := rewardUC.RewardRepo.FindAllRewardHistory(userId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// FindAllUploadReward implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) FindAllUploadReward() ([]entity.UserRewardRequestCore, error) {
	userReward, err := rewardUC.RewardRepo.FindAllUploadReward()
	if err != nil {
		return nil, errors.New("error get user reward request")
	}
	return userReward, nil
}

// UploadRewardRequest implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) UploadRewardRequest(input entity.UserRewardRequestCore) error {
	//tambahin pengecekan point disini nanti
	userData, errUser := rewardUC.UserRepo.ReadSpecificUser(input.UserId)
	if errUser != nil {
		return errors.New("failed get user")
	}

	userPoint, _ := strconv.Atoi(userData.Point)

	rewardData, errReward := rewardUC.RewardRepo.FindById(input.RewardId)
	if errReward != nil {
		return errors.New("failed get reward")
	}

	if userPoint < rewardData.Price{
		return errors.New("not enough point")
	}

	if rewardData.Stock < 1{
		return errors.New("not enough stock")
	}

	err := rewardUC.RewardRepo.UploadRewardRequest(input)
	if err != nil {
		return errors.New("failed request reward")
	}

	return nil
}

// FindUserRewardById implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) FindUserRewardById(id string) (entity.UserRewardRequestCore, error) {
	reward, err := rewardUC.RewardRepo.FindUserRewardById(id)
	if err != nil {
		return entity.UserRewardRequestCore{}, err
	}

	return reward, nil
}

// UpdateReqRewardStatus implements entity.RewardUseCaseInterface.
func (rewardUC *RewardService) UpdateReqRewardStatus(rewardId string, data entity.UserRewardRequestCore) error {
	err := rewardUC.RewardRepo.UpdateReqRewardStatus(rewardId, data)
	if err != nil{
		return err
	}

	return nil
}