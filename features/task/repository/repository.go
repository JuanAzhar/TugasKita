package repository

import (
	"errors"
	"tugaskita/features/task/entity"
	"tugaskita/features/task/model"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) entity.TaskDataInterface {
	return &TaskRepository{
		db: db,
	}
}

// CreateTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) CreateTask(input entity.TaskCore) error {
	data := entity.TaskCoreToTaskModel(input)
	tx := taskRepo.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) DeleteTask(taskId string) error {
	dataTask := model.Task{}

	tx := taskRepo.db.Where("id = ? ", taskId).Delete(&dataTask)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

// FindAllMission implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindAllTask() ([]entity.TaskCore, error) {
	var task []model.Task
	taskRepo.db.Find(&task)
	data := []model.Task{}

	dataTask := entity.ListTaskModelToTaskCore(data)
	return dataTask, nil
}

// FindById implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindById(taskId string) (entity.TaskCore, error) {
	dataTask := model.Task{}

	tx := taskRepo.db.Where("id = ? ", taskId).First(&dataTask)
	if tx.Error != nil {
		return entity.TaskCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.TaskCore{}, errors.New("task not found")
	}

	dataResponse := entity.TaskModelToTaskCore(dataTask)
	return dataResponse, nil
}

// UpdateTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) UpdateTask(taskId string, data entity.TaskCore) error {
	dataTask := entity.TaskCoreToTaskModel(data)
	getData := model.Task{}
	tx := taskRepo.db.Where("id=?", taskId).First(&getData)
	if tx.Error != nil {
		return tx.Error
	}

	tx = taskRepo.db.Where("id = ?", taskId).Updates(&dataTask)
	if tx.Error != nil {
		if tx.Error != nil {
			return tx.Error
		}
		return tx.Error
	}
	return nil
}

// UpdateTaskStatus implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) UpdateTaskStatus(data model.Task) error {
	panic("unimplemented")
}
