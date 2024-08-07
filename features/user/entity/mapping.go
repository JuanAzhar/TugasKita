package entity

import "tugaskita/features/user/model"

func UserCoreToUserModel(data UserCore) model.Users {
	return model.Users{
		ID:         data.ID,
		Name:       data.Name,
		Email:      data.Email,
		Password:   data.Password,
		Role:       data.Role,
		Point:      data.Point,
		TotalPoint: data.TotalPoint,
	}
}

func UserModelToUserCore(data model.Users) UserCore {
	return UserCore{
		ID:         data.ID,
		Name:       data.Name,
		Email:      data.Email,
		Password:   data.Password,
		Role:       data.Role,
		Point:      data.Point,
		TotalPoint: data.TotalPoint,
	}
}
