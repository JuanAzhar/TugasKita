package repository

import (
	"errors"
	"tugaskita/features/reward/entity"
	"tugaskita/features/reward/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) entity.RewardDataInterface {
	return &RewardRepository{
		db: db,
	}
}

// CreateReward implements entity.RewardDataInterface.
func (rewardRepo *RewardRepository) CreateReward(input entity.RewardCore) error {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return UUIDerr
	}

	data := entity.RewardCoreToRewardModel(input)
	data.ID = newUUID
	tx := rewardRepo.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteTask implements entity.RewardDataInterface.
func (rewardRepo *RewardRepository) DeleteReward(rewardId string) error {
	dataReward := model.Reward{}

	tx := rewardRepo.db.Where("id = ? ", rewardId).Delete(&dataReward)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("reward not found")
	}

	return nil
}

// FindAllReward implements entity.RewardDataInterface.
func (rewardRepo *RewardRepository) FindAllReward() ([]entity.RewardCore, error) {
	var reward []model.Reward
	rewardRepo.db.Find(&reward)

	dataReward := entity.ListRewardModelToRewardCore(reward)
	return dataReward, nil
}

// FindById implements entity.RewardDataInterface.
func (rewardRepo *RewardRepository) FindById(rewardId string) (entity.RewardCore, error) {
	dataReward := model.Reward{}

	tx := rewardRepo.db.Where("id = ? ", rewardId).First(&dataReward)
	if tx.Error != nil {
		return entity.RewardCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.RewardCore{}, errors.New("reward not found")
	}

	dataResponse := entity.RewardModelToRewardCore(dataReward)
	return dataResponse, nil
}

// UpdateReward implements entity.RewardDataInterface.
func (rewardRepo *RewardRepository) UpdateReward(rewardId string, data entity.RewardCore) error {
	dataReward := entity.RewardCoreToRewardModel(data)

	tx := rewardRepo.db.Where("id = ?", rewardId).Updates(&dataReward)
	if tx.Error != nil {
		if tx.Error != nil {
			return tx.Error
		}
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("reward not found")
	}

	return nil
}

// FindAllUploadReward implements entity.RewardDataInterface.
func (rewardRepo *RewardRepository) FindAllUploadReward() ([]entity.UserRewardRequestCore, error) {
	var reward []model.UserRewardRequest

	errData := rewardRepo.db.Find(&reward).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]entity.UserRewardRequestCore, len(reward))
	for i, v := range reward {
		mapData[i] = entity.UserRewardRequestCore{
			Id:       v.Id,
			RewardId: v.RewardId,
			UserId:   v.UserId,
			Status:   v.Status,
		}
	}
	return mapData, nil
}

// UploadRewardRequest implements entity.RewardDataInterface.
func (rewardRepo *RewardRepository) UploadRewardRequest(input entity.UserRewardRequestCore) error {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return UUIDerr
	}

	var inputData = model.UserRewardRequest{
		Id:       newUUID,
		RewardId: input.RewardId,
		UserId:   input.UserId,
		Status:   input.Status,
	}

	errUpload := rewardRepo.db.Save(&inputData)
	if errUpload != nil {
		return errUpload.Error
	}

	return nil
}

// FindAllRewardRequestUser implements entity.RewardDataInterface.
func (rewardRepo *RewardRepository) FindAllRewardHistory(userId string) ([]entity.UserRewardRequestCore, error) {
	var reward []model.UserRewardRequest
	rewardRepo.db.Where("user_id=?", userId).Find(&reward)

	dataReward := entity.ListRewardUserModelToListRewardUserCore(reward)
	return dataReward, nil
}

// FindUserRewardById implements entity.RewardDataInterface.
func (rewardRepo *RewardRepository) FindUserRewardById(id string) (entity.UserRewardRequestCore, error) {
	var data model.UserRewardRequest

	errData := rewardRepo.db.Where("id=?", id).First(&data).Error
	if errData != nil {
		return entity.UserRewardRequestCore{}, errData
	}

	userCore := entity.UserRewardRequestCore{
		Id:       data.Id,
		RewardId: data.RewardId,
		UserId:   data.UserId,
		Status:   data.Status,
	}

	return userCore, nil
}
