package repository

import (
	"errors"
	"tugaskita/features/task/entity"
	"tugaskita/features/task/model"

	"github.com/google/uuid"
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
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return UUIDerr
	}

	data := entity.TaskCoreToTaskModel(input)
	data.ID = newUUID
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

	dataTask := entity.ListTaskModelToTaskCore(task)
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

	tx := taskRepo.db.Where("id = ?", taskId).Updates(&dataTask)
	if tx.Error != nil {
		if tx.Error != nil {
			return tx.Error
		}
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

// UpdateTaskStatus implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) UpdateTaskStatus(taskId string, data entity.UserTaskUploadCore) error {
	taskData := entity.TaskUserCoreToTaskUserModel(data)


	tx := taskRepo.db.Where("id=?", taskId).Updates(taskData)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// FindAllClaimedTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindAllClaimedTask(userId string) ([]entity.UserTaskUploadCore, error) {
	var task []model.UserTaskUpload
	taskRepo.db.Where("user_id=?", userId).Find(&task)

	dataTask := entity.ListTaskUserModelToTaskUserCore(task)
	return dataTask, nil
}

// FindAllTaskNotClaimed implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindTasksNotClaimedByUser(userId string) ([]entity.TaskCore, error) {
	var tasks []model.Task
	taskRepo.db.Raw(`
		SELECT * FROM tasks WHERE id NOT IN (
			SELECT task_id FROM user_task_uploads WHERE user_id = ?
		)
	`, userId).Scan(&tasks)

	data := entity.ListTaskModelToTaskCore(tasks)
	return data, nil
}

// UploadTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) UploadTask(input entity.UserTaskUploadCore) error {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return UUIDerr
	}

	var inputData = model.UserTaskUpload{
		Id:          newUUID,
		TaskId:      input.TaskId,
		UserId:      input.UserId,
		Image:       input.Image,
		Description: input.Description,
		Status:      input.Status,
	}

	errUpload := taskRepo.db.Save(&inputData)
	if errUpload != nil {
		return errUpload.Error
	}

	return nil
}
