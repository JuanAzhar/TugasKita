package handler

import (
	"net/http"
	"tugaskita/features/task/dto"
	"tugaskita/features/task/entity"
	user "tugaskita/features/user/entity"
	middleware "tugaskita/utils/jwt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TaskController struct {
	taskUsecase entity.TaskUseCaseInterface
	userUsecase user.UserUseCaseInterface
}

func New(taskUC entity.TaskUseCaseInterface, userUC user.UserUseCaseInterface) *TaskController {
	return &TaskController{
		taskUsecase: taskUC,
		userUsecase: userUC,
	}
}

func (handler *TaskController) AddTask(e echo.Context) error {
	userId, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	input := new(dto.TaskRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.TaskCore{
		AdminId:     userId,
		Title:       input.Title,
		Description: input.Description,
		Point:       input.Point,
		Start_date:  input.Start_date,
		End_date:    input.End_date,
	}

	errTask := handler.taskUsecase.CreateTask(data)
	if errTask != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create task",
			"error":   errTask.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create task",
	})
}

func (handler *TaskController) ReadAllTask(e echo.Context) error {
	userId, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role == "admin" {
		data, err := handler.taskUsecase.FindAllTask()
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]any{
				"message": "error get all task",
			})
		}

		dataList := []dto.TaskResponse{}
		for _, v := range data {
			result := dto.TaskResponse{
				Id:         v.ID.String(),
				Title:      v.Title,
				Point:      v.Point,
				Status:     v.Status,
				Type:       v.Type,
				Start_date: v.Start_date,
				End_date:   v.End_date,
			}
			dataList = append(dataList, result)
		}

		return e.JSON(http.StatusOK, map[string]any{
			"message": "get all admin task",
			"data":    dataList,
		})

	} else if role == "user" {
		data, err := handler.taskUsecase.FindTasksNotClaimedByUser(userId)
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]any{
				"message": "error get all task",
			})
		}

		dataList := []dto.TaskResponse{}
		for _, v := range data {
			result := dto.TaskResponse{
				Id:         v.ID.String(),
				Title:      v.Title,
				Point:      v.Point,
				Status:     v.Status,
				Type:       v.Type,
				Start_date: v.Start_date,
				End_date:   v.End_date,
			}
			dataList = append(dataList, result)
		}

		return e.JSON(http.StatusOK, map[string]any{
			"message": "get all user task",
			"data":    dataList,
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]any{
		"message": "access denied",
	})
}

func (handler *TaskController) ReadSpecificTask(e echo.Context) error {

	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "task not found",
		})
	}

	data, err := handler.taskUsecase.FindById(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific task",
		})
	}

	response := dto.TaskResponseDetail{
		Title:       data.Title,
		Description: data.Description,
		Point:       data.Point,
		Status:      data.Status,
		Start_date:  data.Start_date,
		End_date:    data.End_date,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get task",
		"data":    response,
	})
}

func (handler *TaskController) DeleteTask(e echo.Context) error {
	_, role, errRole := middleware.ExtractTokenUserId(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")
	err := handler.taskUsecase.DeleteTask(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting task",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Task deleted successfully",
	})
}

func (handler *TaskController) UpdateTask(e echo.Context) error {
	adminId, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")

	data := new(dto.TaskRequest)
	if errBind := e.Bind(data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	taskData := entity.TaskCore{
		AdminId:     adminId,
		Title:       data.Title,
		Description: data.Description,
		Point:       data.Point,
		Status:      data.Status,
		Start_date:  data.Start_date,
		End_date:    data.End_date,
	}

	errUpdate := handler.taskUsecase.UpdateTask(idParams, taskData)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating task",
			"error":   errUpdate.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "task updated successfully",
	})
}

func (handler *TaskController) UpdateTaskStatus(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")

	data := dto.UserTaskUploadRequest{}
	if errBind := e.Bind(&data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	//get userId & taskId
	dataTask, _ := handler.taskUsecase.FindUserTaskById(idParams)

	status := entity.UserTaskUploadCore{
		TaskId:      dataTask.TaskId,
		UserId:      dataTask.UserId,
		Image:       data.Image,
		Description: data.Description,
		Status:      data.Status,
		Message:     data.Message,
	}

	errUpdate := handler.taskUsecase.UpdateTaskStatus(idParams, status)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating task status",
			"error":   errUpdate.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "task status updated",
	})
}

func (handler *TaskController) UpdateTaskReqStatus(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")

	data := dto.UserReqTaskRequest{}
	if errBind := e.Bind(&data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	//get userId & taskId
	dataTask, _ := handler.taskUsecase.FindUserTaskReqById(idParams)

	status := entity.UserTaskSubmissionCore{
		UserId:      dataTask.UserId,
		UserName:    dataTask.UserName,
		Title:       data.Title,
		Image:       data.Image,
		Description: data.Description,
		Point:       data.Point,
		Status:      data.Status,
		Message:     data.Message,
	}

	errUpdate := handler.taskUsecase.UpdateTaskReqStatus(idParams, status)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating request task status",
			"error":   errUpdate.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "task request status updated",
	})
}

func (handler *TaskController) ReadHistoryTaskUser(e echo.Context) error {
	userId, _, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	data, err := handler.taskUsecase.FindAllClaimedTask(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get history task",
		})
	}

	dataList := []entity.UserTaskUploadCore{}
	for _, v := range data {

		userData, _ := handler.userUsecase.ReadSpecificUser(v.UserId)
		taskData, _ := handler.taskUsecase.FindById(v.TaskId)

		result := entity.UserTaskUploadCore{
			Id:          v.Id,
			TaskId:      v.TaskId,
			TaskName:    taskData.Title,
			UserId:      v.UserId,
			UserName:    userData.Name,
			Image:       v.Image,
			Description: v.Description,
			Status:      v.Status,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all task history",
		"data":    dataList,
	})
}

func (handler *TaskController) UploadTaskUser(e echo.Context) error {
	input := dto.UserTaskUploadRequest{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "No file uploaded",
			})
		}
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error uploading file",
		})
	}

	userId, _, errRole := middleware.ExtractTokenUserId(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	dataInput := entity.UserTaskUploadCore{
		TaskId:      input.TaskId,
		Image:       input.Image,
		Description: input.Description,
	}
	dataInput.UserId = userId

	errUpload := handler.taskUsecase.UploadTask(dataInput, image)
	if errUpload != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error upload task",
			"error":   errUpload.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes upload task",
		"data":    dataInput,
	})
}

func (handler *TaskController) FindAllUserTask(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	data, err := handler.taskUsecase.FindAllUserTask()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all user task",
		})
	}

	dataList := []entity.UserTaskUploadCore{}
	for _, v := range data {

		userData, _ := handler.userUsecase.ReadSpecificUser(v.UserId)
		taskData, _ := handler.taskUsecase.FindById(v.TaskId)

		result := entity.UserTaskUploadCore{
			Id:          v.Id,
			TaskId:      v.TaskId,
			TaskName:    taskData.Title,
			UserId:      v.UserId,
			UserName:    userData.Name,
			Image:       v.Image,
			Description: v.Description,
			Status:      v.Status,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all user task",
		"data":    dataList,
	})
}

func (handler *TaskController) FindUserTaskById(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")

	data, err := handler.taskUsecase.FindUserTaskById(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific task",
		})
	}

	userData, _ := handler.userUsecase.ReadSpecificUser(data.UserId)
	taskData, _ := handler.taskUsecase.FindById(data.TaskId)

	response := dto.UserTaskUploadResponse{
		Id:          data.Id.String(),
		UserId:      data.UserId,
		UserName:    userData.Name,
		TaskId:      data.TaskId,
		TaskName:    taskData.Title,
		Image:       data.Image,
		Description: data.Description,
		Status:      data.Status,
		Message:     data.Message,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get task",
		"data":    response,
	})
}

func (handler *TaskController) FindUserTaskReqyId(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")

	data, err := handler.taskUsecase.FindUserTaskReqById(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific task",
		})
	}

	userData, _ := handler.userUsecase.ReadSpecificUser(data.UserId)

	response := dto.UserReqTaksResponse{
		Id:          data.Id.String(),
		Title:       data.Title,
		UserId:      data.UserId,
		UserName:    userData.Name,
		Point:       data.Point,
		Image:       data.Image,
		Description: data.Description,
		Status:      data.Status,
		Message:     data.Message,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get task",
		"data":    response,
	})
}

func (handler *TaskController) UploadRequestTaskUser(e echo.Context) error {
	input := dto.UserReqTaskRequest{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	userId, _, errRole := middleware.ExtractTokenUserId(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "No file uploaded",
			})
		}
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error uploading file",
		})
	}

	dataInput := entity.UserTaskSubmissionCore{
		Title:       input.Title,
		Point:       input.Point,
		Image:       input.Image,
		Description: input.Description,
	}
	dataInput.UserId = userId

	errUpload := handler.taskUsecase.UploadTaskRequest(dataInput, image)
	if errUpload != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error upload request task",
			"error":   errUpload.Error(),
		})
	}

	dataRespon := dto.UserReqTaksResponse{
		Title:       input.Title,
		Image:       input.Image,
		Description: input.Description,
		Point:       input.Point,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes upload request task",
		"data":    dataRespon,
	})
}

func (handler *TaskController) FindAllUserRequestTask(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	data, err := handler.taskUsecase.FindAllRequestTask()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all user request task",
		})
	}

	dataList := []entity.UserTaskSubmissionCore{}
	for _, v := range data {

		userData, _ := handler.userUsecase.ReadSpecificUser(v.UserId)

		result := entity.UserTaskSubmissionCore{
			Id:          v.Id,
			Title:       v.Title,
			Point:       v.Point,
			UserId:      v.UserId,
			UserName:    userData.Name,
			Image:       v.Image,
			Description: v.Description,
			Status:      v.Status,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all user task request",
		"data":    dataList,
	})
}

func (handler *TaskController) FindAllRequestTaskHistory(e echo.Context) error {
	userId, _, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	data, err := handler.taskUsecase.FindAllRequestTaskHistory(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all user request task history",
		})
	}

	dataList := []entity.UserTaskSubmissionCore{}
	for _, v := range data {
		result := entity.UserTaskSubmissionCore{
			Id:          v.Id,
			Title:       v.Title,
			Point:       v.Point,
			UserId:      v.UserId,
			Image:       v.Image,
			Description: v.Description,
			Status:      v.Status,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all user task request history",
		"data":    dataList,
	})
}
