package service

import (
	"errors"
	"time"
	"tugaskita/features/task/entity"
	"tugaskita/features/task/model"
)

type taskService struct {
	TaskRepo entity.TaskDataInterface
}

func NewTaskService(taskRepo entity.TaskDataInterface) entity.TaskUseCaseInterface {
	return &taskService{
		TaskRepo: taskRepo,
	}
}

// CreateTask implements entity.TaskUseCaseInterface.
func (taskUC *taskService) CreateTask(input entity.TaskCore) error {
	layout := "2006-01-02"
	currentTime := time.Now().Truncate(24 * time.Hour)

	if input.Title == "" || input.Description == "" {
		return errors.New("title and description can't empty")
	}

	start, errStart := time.Parse(layout, input.Start_date.String())
	if errStart != nil {
		return errors.New("start date must be in 'yyyy-mm-dd' format")
	}
	if start.Before(currentTime) {
		return errors.New("please choose at least today")
	}

	end, errEnd := time.Parse(layout, input.End_date.String())
	if errEnd != nil {
		return errors.New("end date must be in 'yyyy-mm-dd' format")
	}

	if end.Before(start) {
		return errors.New("end date must be after start date")
	}

	if end.Equal(start) {
		return errors.New("end date must be different from start date")
	}

	if input.Point <= 0{
		return errors.New("point must be more than 0")
	}

	err := taskUC.TaskRepo.CreateTask(input)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTask implements entity.TaskCoreUseCaseInterface.
func (taskUC *taskService) DeleteTask(taskId string) error {
	if taskId == "" {
		return errors.New("insert user id")
	}

	_, err := taskUC.TaskRepo.FindById(taskId)
	if err != nil {
		return errors.New("task not found")
	}

	errDelete := taskUC.TaskRepo.DeleteTask(taskId)
	if errDelete != nil {
		return errors.New("can't delete task")
	}

	return nil
}

// FindAllMission implements entity.TaskCoreUseCaseInterface.
func (taskUC *taskService) FindAllTask() ([]entity.TaskCore, error) {
	data, err := taskUC.TaskRepo.FindAllTask()
	if err != nil {
		return nil, err
	}

	return data, nil

}

// FindById implements entity.TaskCoreUseCaseInterface.
func (taskUC *taskService) FindById(taskId string) (entity.TaskCore, error) {
	if taskId == "" {
		return entity.TaskCore{}, errors.New("task ID is required")
	}

	task, err := taskUC.TaskRepo.FindById(taskId)
	if err != nil {
		return entity.TaskCore{}, err
	}

	return task, nil
}

// UpdateTask implements entity.TaskCoreUseCaseInterface.
func (taskUC *taskService) UpdateTask(taskId string, data entity.TaskCore) error {
	panic("unimplemented")
}

// UpdateTaskStatus implements entity.TaskUseCaseInterface.
func (taskUC *taskService) UpdateTaskStatus(data model.Task) error {
	panic("unimplemented")
}
