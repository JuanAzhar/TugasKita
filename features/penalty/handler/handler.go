package handler

import (
	"net/http"
	"tugaskita/features/penalty/dto"
	"tugaskita/features/penalty/entity"
	middleware "tugaskita/utils/jwt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PenaltyController struct {
	penaltyUsecase entity.PenaltyUseCaseInterface
}

func New(penaltyUC entity.PenaltyUseCaseInterface) *PenaltyController {
	return &PenaltyController{
		penaltyUsecase: penaltyUC,
	}
}

func (handler *PenaltyController) CreatePenalty(e echo.Context) error {
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

	input := dto.PenaltyRequest{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.PenaltyCore{
		UserId:      input.UserId,
		Description: input.Description,
		Point:       input.Point,
		Date:        input.Date,
	}

	errTask := handler.penaltyUsecase.CreatePenalty(data)
	if errTask != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create penalty",
			"error":   errTask.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create penalty",
	})
}

func (handler *PenaltyController) DeletePenalty(e echo.Context) error {
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
	err := handler.penaltyUsecase.DeletePenalty(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting penalty",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "penalty deleted successfully",
	})
}

func (handler *PenaltyController) FindAllPenalty(e echo.Context) error {
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

	data, err := handler.penaltyUsecase.FindAllPenalty()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all penalty user",
		})
	}

	dataList := []entity.PenaltyCore{}
	for _, v := range data {
		result := entity.PenaltyCore{
			Id:          v.Id,
			UserId:      v.UserId,
			Description: v.Description,
			Point:       v.Point,
			Date:        v.Date,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all penalty user",
		"data":    dataList,
	})
}

func (handler *PenaltyController) FindSpecificPenalty(e echo.Context) error {
	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "reward not found",
		})
	}

	data, err := handler.penaltyUsecase.FindSpecificPenalty(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific penalty",
		})
	}

	response := entity.PenaltyCore{
		Id:          data.Id,
		UserId:      data.UserId,
		Description: data.Description,
		Point:       data.Point,
		Date:        data.Date,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get specific penalty",
		"data":    response,
	})
}

func (handler *PenaltyController) UpdatePenalty(e echo.Context) error {
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

	data := new(dto.PenaltyRequest)
	if errBind := e.Bind(data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	rewardData := entity.PenaltyCore{
		UserId:      data.UserId,
		Description: data.Description,
		Point:       data.Point,
		Date:        data.Date,
	}

	errUpdate := handler.penaltyUsecase.UpdatePenalty(idParams, rewardData)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating penalty",
			"error":   errUpdate.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "reward penalty successfully",
	})
}

func (handler *PenaltyController) FindAllPenaltyHistory(e echo.Context) error {
	userId, _, errRole := middleware.ExtractTokenUserId(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	data, err := handler.penaltyUsecase.FindAllPenaltyHistory(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all penalty history",
		})
	}

	dataList := []entity.PenaltyCore{}
	for _, v := range data {
		result := entity.PenaltyCore{
			Id:          v.Id,
			UserId:      v.UserId,
			Description: v.Description,
			Point:       v.Point,
			Date:        v.Date,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all penalty history",
		"data":    dataList,
	})
}