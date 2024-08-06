package service

import (
	"errors"
	"time"
	"tugaskita/features/task/entity"
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
func (taskUC *taskService) CreateTask(data entity.TaskCore) error {
	layout := "2006-01-02"
	currentTime := time.Now().Truncate(24 * time.Hour)

	if data.Title == "" || data.Description == "" {
		return errors.New("title and description can't empty")
	}

	start, errStart := time.Parse(layout, data.Start_date)
	if errStart != nil {
		return errors.New("start date must be in 'yyyy-mm-dd' format")
	}
	if start.Before(currentTime) {
		return errors.New("please choose at least today")
	}

	end, errEnd := time.Parse(layout, data.End_date)
	if errEnd != nil {
		return errors.New("end date must be in 'yyyy-mm-dd' format")
	}

	if end.Before(start) {
		return errors.New("end date must be after start date")
	}

	if end.Equal(start) {
		return errors.New("end date must be different from start date")
	}

	if data.Point <= 0 {
		return errors.New("point must be more than 0")
	}

	err := taskUC.TaskRepo.CreateTask(data)
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
		return nil, errors.New("error get data")
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
	layout := "2006-01-02"
	currentTime := time.Now().Truncate(24 * time.Hour)

	if data.Title == "" || data.Description == "" {
		return errors.New("title and description can't empty")
	}

	start, errStart := time.Parse(layout, data.Start_date)
	if errStart != nil {
		return errors.New("start date must be in 'yyyy-mm-dd' format")
	}
	if start.Before(currentTime) {
		return errors.New("please choose at least today")
	}

	end, errEnd := time.Parse(layout, data.End_date)
	if errEnd != nil {
		return errors.New("end date must be in 'yyyy-mm-dd' format")
	}

	if end.Before(start) {
		return errors.New("end date must be after start date")
	}

	if end.Equal(start) {
		return errors.New("end date must be different from start date")
	}

	if data.Point <= 0 {
		return errors.New("point must be more than 0")
	}

	err := taskUC.TaskRepo.UpdateTask(taskId, data)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTaskStatus implements entity.TaskUseCaseInterface.
func (taskUC *taskService) UpdateTaskStatus(taskId string, data entity.UserTaskUploadCore) error {
	err := taskUC.TaskRepo.UpdateTaskStatus(taskId, data)
	if err != nil {
		return err
	}

	return nil
}

// FindAllClaimedTask implements entity.TaskUseCaseInterface.
func (taskUC *taskService) FindAllClaimedTask(userId string) ([]entity.UserTaskUploadCore, error) {
	data, err := taskUC.TaskRepo.FindAllClaimedTask(userId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// FindTasksNotClaimedByUser implements entity.TaskUseCaseInterface.
func (taskUC *taskService) FindTasksNotClaimedByUser(userId string) ([]entity.TaskCore, error) {
	data, err := taskUC.TaskRepo.FindTasksNotClaimedByUser(userId)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UploadTask implements entity.TaskUseCaseInterface.
func (taskUC *taskService) UploadTask(data entity.UserTaskUploadCore) error {
	_, errTask := taskUC.TaskRepo.FindById(data.TaskId)
	if errTask != nil {
		return errors.New("task not found")
	}

	if data.Image == "" || data.Description == "" {
		return errors.New("image and description can't empty")
	}

	err := taskUC.TaskRepo.UploadTask(data)
	if err != nil {
		return errors.New("failed upload task")
	}

	return nil
}

// FindUserTask implements entity.TaskUseCaseInterface.
func (taskUC *taskService) FindAllUserTask() ([]entity.UserTaskUploadCore, error) {
	userTask, err := taskUC.TaskRepo.FindAllUserTask()
	if err != nil {
		return nil, errors.New("error get user task")
	}

	return userTask, nil
}

// FindUserTaskById implements entity.TaskUseCaseInterface.
func (taskUC *taskService) FindUserTaskById(id string) (entity.UserTaskUploadCore, error) {
	task, err := taskUC.TaskRepo.FindUserTaskById(id)
	if err != nil {
		return entity.UserTaskUploadCore{}, err
	}

	return task, nil
}

// UploadTaskRequest implements entity.TaskUseCaseInterface.
func (taskUC *taskService) UploadTaskRequest(input entity.UserTaskSubmissionCore) error {
	if input.Image == "" || input.Description == "" || input.Title == "" {
		return errors.New("image, description and title can't empty")
	}

	if input.Point <= 0 {
		return errors.New("point can't less then 0")
	}

	err := taskUC.TaskRepo.UploadTaskRequest(input)
	if err != nil {
		return errors.New("failed upload request task")
	}

	return nil
}

// FindAllRequestTask implements entity.TaskUseCaseInterface.
func (taskUC *taskService) FindAllRequestTask() ([]entity.UserTaskSubmissionCore, error) {
	userTask, err := taskUC.TaskRepo.FindAllRequestTask()
	if err != nil {
		return nil, errors.New("error get user request task")
	}

	return userTask, nil
}