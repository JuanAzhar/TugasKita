package repository

import (
	"errors"
	"mime/multipart"
	"strconv"
	"time"
	"tugaskita/features/task/entity"
	"tugaskita/features/task/model"
	user "tugaskita/features/user/entity"
	userModel "tugaskita/features/user/model"
	"tugaskita/utils/cloudinary"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db             *gorm.DB
	userRepository user.UserDataInterface
}

func NewTaskRepository(db *gorm.DB, userRepository user.UserDataInterface) entity.TaskDataInterface {
	return &TaskRepository{
		db:             db,
		userRepository: userRepository,
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
	var pointTask model.Task
	var userData userModel.Users
	taskData := entity.TaskUserCoreToTaskUserModel(data)

	// get task data
	errData := taskRepo.db.Where("id=?", data.TaskId).First(&pointTask).Error
	if errData != nil {
		return errData
	}

	// get user
	errUser := taskRepo.db.Where("id=?", data.UserId).First(&userData).Error
	if errUser != nil {
		return errUser
	}

	// update status
	tx := taskRepo.db.Where("id=?", taskId).Updates(taskData)
	if tx.Error != nil {
		return tx.Error
	}

	if taskData.Status == "Diterima" {
		userPoint, _ := strconv.Atoi(userData.Point)
		count := userPoint + pointTask.Point

		userData.Point = strconv.Itoa(count)

		saveUser := user.UserModelToUserCore(userData)

		updateUser := taskRepo.userRepository.UpdatePoint(data.UserId, saveUser)
		if updateUser != nil {
			return updateUser
		}
	}

	return nil
}

// UpdateTaskReqStatus implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) UpdateTaskReqStatus(id string, data entity.UserTaskSubmissionCore) error {
	var pointTask model.UserTaskSubmission
	var userData userModel.Users
	taskData := entity.TaskUserReqCoreToTaskUserReqModel(data)

	// get user
	errUser := taskRepo.db.Where("id=?", data.UserId).First(&userData).Error
	if errUser != nil {
		return errUser
	}

	// update status
	tx := taskRepo.db.Where("id=?", id).Updates(taskData)
	if tx.Error != nil {
		return tx.Error
	}

	// get task data
	errData := taskRepo.db.Where("id=?", id).First(&pointTask).Error
	if errData != nil {
		return errData
	}

	if taskData.Status == "Diterima" {
		userPoint, _ := strconv.Atoi(userData.Point)
		count := userPoint + pointTask.Point

		userData.Point = strconv.Itoa(count)

		saveUser := user.UserModelToUserCore(userData)

		updateUser := taskRepo.userRepository.UpdatePoint(data.UserId, saveUser)
		if updateUser != nil {
			return updateUser
		}
	}

	return nil
}

// FindAllClaimedTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindAllClaimedTask(userId string) ([]entity.UserTaskUploadCore, error) {
	var task []model.UserTaskUpload
	taskRepo.db.Where("user_id=?", userId).Find(&task)

	dataTask := make([]entity.UserTaskUploadCore, len(task))
	for i, v := range task {
		dataTask[i] = entity.UserTaskUploadCore{
			Id:          v.Id,
			TaskId:      v.TaskId,
			UserId:      v.UserId,
			Image:       v.Image,
			Description: v.Description,
			Status:      v.Status,
			Message:     v.Message,
		}
	}
	return dataTask, nil
}

// FindAllRequestTaskHistory implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindAllRequestTaskHistory(userId string) ([]entity.UserTaskSubmissionCore, error) {
	var task []model.UserTaskSubmission
	taskRepo.db.Where("user_id=?", userId).Find(&task)

	dataTask := entity.ListTaskUserReqModelToTaskUserReqCore(task)
	return dataTask, nil
}

// FindAllTaskNotClaimed implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindTasksNotClaimedByUser(userId string) ([]entity.TaskCore, error) {
	var tasks []model.Task

	currentDate := time.Now().Format("2006-01-02")

	taskRepo.db.Raw(`
	SELECT * FROM tasks 
	WHERE id NOT IN (
		SELECT task_id FROM user_task_uploads 
		WHERE user_id = ? AND status != 'Ditolak'
	) AND status = 'Active' AND end_date >= ?
`, userId, currentDate).Scan(&tasks)

	data := entity.ListTaskModelToTaskCore(tasks)
	return data, nil
}

// UploadTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) UploadTask(input entity.UserTaskUploadCore, image *multipart.FileHeader) error {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return UUIDerr
	}

	file, err := image.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	imageURL, err := cloudinary.UploadToCloudinary(file, image.Filename)
	if err != nil {
		return err
	}

	input.Image = imageURL

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

// FindUserTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindAllUserTask() ([]entity.UserTaskUploadCore, error) {
	var userTask []model.UserTaskUpload

	errData := taskRepo.db.Find(&userTask).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]entity.UserTaskUploadCore, len(userTask))
	for i, v := range userTask {
		mapData[i] = entity.UserTaskUploadCore{
			Id:          v.Id,
			TaskId:      v.TaskId,
			UserId:      v.UserId,
			Image:       v.Image,
			Description: v.Description,
			Status:      v.Status,
			Message:     v.Message,
		}
	}
	return mapData, nil
}

// FindUserTaskById implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindUserTaskById(id string) (entity.UserTaskUploadCore, error) {
	var data model.UserTaskUpload

	errData := taskRepo.db.Where("id=?", id).First(&data).Error
	if errData != nil {
		return entity.UserTaskUploadCore{}, errData
	}

	userData, _ := taskRepo.userRepository.ReadSpecificUser(data.UserId)
	taskData, _ := taskRepo.FindById(data.TaskId)

	userCore := entity.UserTaskUploadCore{
		Id:          data.Id,
		TaskId:      data.TaskId,
		TaskName:    taskData.Title,
		UserId:      data.UserId,
		UserName:    userData.Name,
		Image:       data.Image,
		Description: data.Description,
		Status:      data.Status,
	}

	return userCore, nil
}

// FindUserTaskReqById implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindUserTaskReqById(id string) (entity.UserTaskSubmissionCore, error) {
	var data model.UserTaskSubmission

	errData := taskRepo.db.Where("id=?", id).First(&data).Error
	if errData != nil {
		return entity.UserTaskSubmissionCore{}, errData
	}

	userData, _ := taskRepo.userRepository.ReadSpecificUser(data.UserId)

	userCore := entity.UserTaskSubmissionCore{
		Id:          data.Id,
		Title:       data.Title,
		UserId:      data.UserId,
		UserName:    userData.Name,
		Image:       data.Image,
		Point:       data.Point,
		Description: data.Description,
		Status:      data.Status,
		Message:     data.Message,
	}

	return userCore, nil
}

// UploadTaskRequest implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) UploadTaskRequest(input entity.UserTaskSubmissionCore, image *multipart.FileHeader) error {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return UUIDerr
	}

	file, err := image.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	imageURL, err := cloudinary.UploadToCloudinary(file, image.Filename)
	if err != nil {
		return err
	}

	input.Image = imageURL

	var inputData = model.UserTaskSubmission{
		Id:          newUUID,
		UserId:      input.UserId,
		Title:       input.Title,
		Point:       input.Point,
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

// FindAllRequestTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) FindAllRequestTask() ([]entity.UserTaskSubmissionCore, error) {
	var userTask []model.UserTaskSubmission

	errData := taskRepo.db.Find(&userTask).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]entity.UserTaskSubmissionCore, len(userTask))
	for i, v := range userTask {
		mapData[i] = entity.UserTaskSubmissionCore{
			Id:          v.Id,
			UserId:      v.UserId,
			Title:       v.Title,
			Image:       v.Image,
			Description: v.Description,
			Point:       v.Point,
			Status:      v.Status,
		}
	}
	return mapData, nil
}

// CountUserClearTask implements entity.TaskDataInterface.
func (taskRepo *TaskRepository) CountUserClearTask(id string) (int, error) {
	var countUpload int64
	var countSubmission int64

	// Menghitung jumlah data pada tabel UserTaskUpload dengan status 'Diterima'
	errUpload := taskRepo.db.Model(&model.UserTaskUpload{}).
		Where("user_id = ? AND status = ?", id, "Diterima").
		Count(&countUpload).Error
	if errUpload != nil {
		return 0, errUpload
	}

	// Menghitung jumlah data pada tabel UserTaskSubmission dengan status 'Diterima'
	errSubmission := taskRepo.db.Model(&model.UserTaskSubmission{}).
		Where("user_id = ? AND status = ?", id, "Diterima").
		Count(&countSubmission).Error
	if errSubmission != nil {
		return 0, errSubmission
	}

	// Menggabungkan hasil dari kedua tabel
	totalCount := int(countUpload + countSubmission)

	return totalCount, nil
}