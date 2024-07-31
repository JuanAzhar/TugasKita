package entity

import (
	"time"
	"tugaskita/features/task/model"
)

func TaskCoreToTaskModel(data TaskCore) model.Task {
	return model.Task{
		ID:          data.ID,
		AdminId:     data.AdminId,
		Title:       data.Title,
		Description: data.Description,
		Point:       data.Point,
		Message:     data.Message,
		Status:      data.Status,
		Type:        data.Type,
		Start_date:  data.Start_date,
		End_date:    data.End_date,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

}

func TaskModelToTaskCore(data model.Task) TaskCore {
	return TaskCore{
		ID:          data.ID,
		AdminId:     data.AdminId,
		Title:       data.Title,
		Description: data.Description,
		Point:       data.Point,
		Message:     data.Message,
		Status:      data.Status,
		Type:        data.Type,
		Start_date:  data.Start_date,
		End_date:    data.End_date,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

}

func ListTaskModelToTaskCore(data []model.Task) []TaskCore {
	dataTask := []TaskCore{}
	for _, v := range data {
		result := TaskModelToTaskCore(v)
		dataTask = append(dataTask, result)
	}
	return dataTask
}

func TaskUserModelToTaskUserCore(data model.UserTaskUpload) UserTaskUploadCore {
	return UserTaskUploadCore{
		Id:          data.Id,
		TaskId:      data.TaskId,
		UserId:      data.UserId,
		Image:       data.Image,
		Description: data.Description,
		Status:      data.Status,
	}

}

func TaskUserCoreToTaskUserModel(data UserTaskUploadCore) model.UserTaskUpload {
	return model.UserTaskUpload{
		Id:          data.Id,
		TaskId:      data.TaskId,
		UserId:      data.UserId,
		Image:       data.Image,
		Description: data.Description,
		Status:      data.Status,
	}
}

func ListTaskUserModelToTaskUserCore(data []model.UserTaskUpload) []UserTaskUploadCore {
	dataTask := []UserTaskUploadCore{}
	for _, v := range data {
		result := TaskUserModelToTaskUserCore(v)
		dataTask = append(dataTask, result)
	}
	return dataTask
}
