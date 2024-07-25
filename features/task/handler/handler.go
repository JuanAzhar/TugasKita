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
	data, err := handler.taskUsecase.FindAllTask()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all task",
		})
	}

	dataList := []dto.TaskRequest{}
	for _, v := range data {
		result := dto.TaskRequest{
			Title:      v.Title,
			Point:      v.Point,
			Start_date: v.Start_date,
			End_date:   v.End_date,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all user",
		"data":    dataList,
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
