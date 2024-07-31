package handler

import (
	"net/http"
	"tugaskita/features/task/dto"
	"tugaskita/features/task/entity"
	middleware "tugaskita/utils/jwt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TaskController struct {
	taskUsecase entity.TaskUseCaseInterface
}

func New(taskUC entity.TaskUseCaseInterface) *TaskController {
	return &TaskController{
		taskUsecase: taskUC,
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
		"message": "succes create user",
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
		"message": "get user",
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

	status := entity.UserTaskUploadCore{
		TaskId:      data.TaskId,
		UserId:      data.UserId,
		Image:       data.Image,
		Description: data.Description,
		Status:      data.Status,
	}

	errUpdate := handler.taskUsecase.UpdateTaskStatus(idParams, status)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating task status",
			"error":   errUpdate.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "task updated successfully",
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

	dataList := []dto.UserTaskUploadResponse{}
	for _, v := range data {
		result := dto.UserTaskUploadResponse{
			Id:          v.Id.String(),
			TaskId:      v.TaskId,
			UserId:      v.UserId,
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
	input := new(dto.UserTaskUploadRequest)
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

	println("ini id task", input.TaskId)

	dataInput := entity.UserTaskUploadCore{
		TaskId:      input.TaskId,
		Image:       input.Image,
		Description: input.Description,
	}
	dataInput.UserId = userId

	err := handler.taskUsecase.UploadTask(dataInput)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error upload task",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes upload task",
		"data":    dataInput,
	})
}