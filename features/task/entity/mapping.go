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
		Message:     data.Message,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
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
		Message:     data.Message,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
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

func TaskUserReqModelToTaskUserReqCore(data model.UserTaskSubmission) UserTaskSubmissionCore {
	return UserTaskSubmissionCore{
		Id:          data.Id,
		Title:       data.Title,
		Point:       data.Point,
		UserId:      data.UserId,
		Image:       data.Image,
		Description: data.Description,
		Status:      data.Status,
		Message:     data.Message,
		Type:        data.Type,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func TaskUserReqCoreToTaskUserReqModel(data UserTaskSubmissionCore) model.UserTaskSubmission {
	return model.UserTaskSubmission{
		Id:          data.Id,
		Title:       data.Title,
		Point:       data.Point,
		UserId:      data.UserId,
		Image:       data.Image,
		Description: data.Description,
		Status:      data.Status,
		Message:     data.Message,
		Type:        data.Type,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func ListTaskUserReqModelToTaskUserReqCore(data []model.UserTaskSubmission) []UserTaskSubmissionCore {
	dataTask := []UserTaskSubmissionCore{}
	for _, v := range data {
		result := TaskUserReqModelToTaskUserReqCore(v)
		dataTask = append(dataTask, result)
	}
	return dataTask
}
