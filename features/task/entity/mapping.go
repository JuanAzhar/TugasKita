package entity

import (
	"time"
	"tugaskita/features/task/model"
)

func TaskCoreToTaskModel(data TaskCore) model.Task {
	return model.Task{
		AdminId:     data.AdminId,
		Title:       data.Title,
		Description: data.Description,
		Point:       data.Point,
		Message:     data.Message,
		Status:      data.Status,
		Start_date:  data.Start_date,
		End_date:    data.End_date,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

}

func TaskModelToTaskCore(data model.Task) TaskCore {
	return TaskCore{
		AdminId:     data.AdminId,
		Title:       data.Title,
		Description: data.Description,
		Point:       data.Point,
		Message:     data.Message,
		Status:      data.Status,
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