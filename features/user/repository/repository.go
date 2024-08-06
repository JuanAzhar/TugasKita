package repository

import (
	"tugaskita/features/user/entity"
	"tugaskita/features/user/model"
	bcrypt "tugaskita/utils/bcrypt"
	utils "tugaskita/utils/jwt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) entity.UserDataInterface {
	return &userRepository{
		db: db,
	}
}

// DeleteUser implements entity.UserDataInterface.
func (userRepo *userRepository) DeleteUser(id string) (err error) {
	var chekcId model.Users

	errData := userRepo.db.Where("id = ?", id).Delete(&chekcId)
	if errData != nil {
		return errData.Error
	}

	return nil

}

// Login implements entity.UserDataInterface.
func (userRepo *userRepository) Login(email string, password string) (entity.UserCore, string, error) {
	var data model.Users

	bcrypt.CheckPasswordHash(data.Password, password)

	tx := userRepo.db.Where("email=?", email).First(&data)
	if tx.Error != nil {
		return entity.UserCore{}, "", tx.Error
	}

	var token string

	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = utils.CreateToken(data.ID, data.Role)
		if errToken != nil {
			return entity.UserCore{}, "", errToken
		}
	}

	var resp = entity.UserCore{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	return resp, token, nil
}

// ReadSpecificUser implements entity.UserDataInterface.
func (userRepo *userRepository) ReadSpecificUser(id string) (user entity.UserCore, err error) {
	var data model.Users
	errData := userRepo.db.Where("id=?", id).First(&data).Error
	if errData != nil {
		return entity.UserCore{}, errData
	}

	userCore := entity.UserCore{
		ID:         data.ID,
		Name:       data.Name,
		Email:      data.Email,
		Point:      data.Point,
		TotalPoint: data.TotalPoint,
		Role:       data.Role,
	}

	return userCore, nil
}

// Register implements entity.UserDataInterface.
func (userRepo *userRepository) Register(data entity.UserCore) (row int, err error) {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return 0, UUIDerr
	}

	hashPassword, err := bcrypt.HashPassword(data.Password)
	if err != nil {
		return 0, err
	}

	var input = model.Users{
		ID:       newUUID.String(),
		Name:     data.Name,
		Email:    data.Email,
		Password: hashPassword,
		Point:    "0",
		Role:     "user",
	}

	erruser := userRepo.db.Save(&input)
	if erruser.Error != nil {
		return 0, erruser.Error
	}

	return 1, nil
}

// ReadAllUser implements entity.UserDataInterface.
func (userRepo *userRepository) ReadAllUser() ([]entity.UserCore, error) {
	var dataUser []model.Users

	errData := userRepo.db.Find(&dataUser).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]entity.UserCore, len(dataUser))
	for i, value := range dataUser {
		mapData[i] = entity.UserCore{
			ID:         value.ID,
			Name:       value.Name,
			Email:      value.Email,
			Role:       value.Role,
			Point:      value.Point,
			TotalPoint: value.TotalPoint,
		}
	}
	return mapData, nil
}

// UpdatePoint implements entity.UserDataInterface.
func (userRepo *userRepository) UpdatePoint(id string, data entity.UserCore) error {
	userData := entity.UserCoreToUserModel(data)

	tx := userRepo.db.Where("id = ?", id).Updates(&userData)
	if tx != nil {
		return tx.Error
	}

	return nil
}

// GetRankUser implements entity.UserDataInterface.
func (userRepo *userRepository) GetRankUser() ([]entity.UserCore, error) {
	var dataUser []model.Users

	errData := userRepo.db.Order("point desc").Find(&dataUser).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]entity.UserCore, len(dataUser))
	for i, value := range dataUser {
		mapData[i] = entity.UserCore{
			ID:         value.ID,
			Name:       value.Name,
			Email:      value.Email,
			Role:       value.Role,
			Point:      value.Point,
			TotalPoint: value.TotalPoint,
		}
	}
	return mapData, nil
}
