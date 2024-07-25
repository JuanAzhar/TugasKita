package entity

import "tugaskita/features/task/model"

type TaskDataInterface interface {
	CreateTask(input TaskCore) error
	FindAllTask() ([]TaskCore, error)
	FindById(taskId string) (TaskCore, error)
	UpdateTask(taskId string, data TaskCore) error
	DeleteTask(taskId string) error
	UpdateTaskStatus(data model.Task) error
}

type TaskUseCaseInterface interface {
	CreateTask(input TaskCore) error
	FindAllTask() ([]TaskCore, error)
	FindById(taskId string) (TaskCore, error)
	UpdateTask(taskId string, data TaskCore) error
	DeleteTask(taskId string) error
	UpdateTaskStatus(data model.Task) error
}
